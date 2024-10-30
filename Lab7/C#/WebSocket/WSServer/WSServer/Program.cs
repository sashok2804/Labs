using System;
using System.Collections.Concurrent;
using System.Net;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace WebSocketChatServer
{
	class Program
	{
		private static readonly HttpListener HttpListener = new HttpListener();
		private static readonly ConcurrentDictionary<WebSocket, bool> Clients = new ConcurrentDictionary<WebSocket, bool>();

		static async Task Main(string[] args)
		{
			HttpListener.Prefixes.Add("http://localhost:8080/");
			HttpListener.Start();
			Console.WriteLine("WebSocket сервер запущен на ws://localhost:8080/");

			while (true)
			{
				var httpContext = await HttpListener.GetContextAsync();
				if (httpContext.Request.IsWebSocketRequest)
				{
					var webSocketContext = await httpContext.AcceptWebSocketAsync(null);
					var webSocket = webSocketContext.WebSocket;
					Clients.TryAdd(webSocket, true);
					Console.WriteLine("Клиент подключился.");

					Task.Run(() => HandleClientAsync(webSocket));
				}
				else
				{
					httpContext.Response.StatusCode = 400;
					httpContext.Response.Close();
				}
			}
		}

		private static async Task HandleClientAsync(WebSocket webSocket)
		{
			var buffer = new byte[1024 * 4];
			try
			{
				while (webSocket.State == WebSocketState.Open)
				{
					var result = await webSocket.ReceiveAsync(new ArraySegment<byte>(buffer), CancellationToken.None);

					if (result.MessageType == WebSocketMessageType.Close)
					{
						await DisconnectClientAsync(webSocket, "Клиент закрыл соединение");
						break;
					}

					var message = Encoding.UTF8.GetString(buffer, 0, result.Count);
					Console.WriteLine($"Получено сообщение: {message}");
					await BroadcastMessageAsync(message);
				}
			}
			catch (WebSocketException ex)
			{
				Console.WriteLine($"Ошибка WebSocket: {ex.Message}");
				await DisconnectClientAsync(webSocket, "Ошибка WebSocket");
			}
		}

		private static async Task DisconnectClientAsync(WebSocket webSocket, string reason)
		{
			try
			{
				if (webSocket.State == WebSocketState.Open)
				{
					await webSocket.CloseAsync(WebSocketCloseStatus.NormalClosure, reason, CancellationToken.None);
				}
			}
			catch (Exception ex)
			{
				Console.WriteLine($"Ошибка при закрытии подключения: {ex.Message}");
			}
			finally
			{
				Clients.TryRemove(webSocket, out _);
				Console.WriteLine("Клиент отключился: " + reason);
			}
		}

		private static async Task BroadcastMessageAsync(string message)
		{
			var messageBuffer = Encoding.UTF8.GetBytes(message);
			var disconnectedClients = new List<WebSocket>();

			foreach (var client in Clients.Keys)
			{
				if (client.State == WebSocketState.Open)
				{
					try
					{
						await client.SendAsync(new ArraySegment<byte>(messageBuffer), WebSocketMessageType.Text, true, CancellationToken.None);
					}
					catch (WebSocketException)
					{
						Console.WriteLine("Ошибка при отправке сообщения клиенту. Отключение клиента.");
						disconnectedClients.Add(client);
					}
				}
				else
				{
					disconnectedClients.Add(client);
				}
			}

			// Удаляем отключённых клиентов из списка
			foreach (var client in disconnectedClients)
			{
				await DisconnectClientAsync(client, "Отключение из-за ошибки при отправке");
			}
		}
	}
}
