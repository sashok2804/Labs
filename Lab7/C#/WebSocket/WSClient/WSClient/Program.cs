using System;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace WebSocketChatClient
{
	class Program
	{
		private static ClientWebSocket _webSocket = new ClientWebSocket();

		static async Task Main(string[] args)
		{
			await ConnectToServerAsync("ws://localhost:8080/");
			Console.WriteLine("Вы подключены к серверу чата. Введите сообщения для отправки.");

			var receiveTask = ReceiveMessagesAsync();

			// Ввод и отправка сообщений пользователем
			while (true)
			{
				var message = Console.ReadLine();

				if (message?.ToLower() == "exit")
				{
					Console.WriteLine("Вы отключены от сервера.");
					await _webSocket.CloseAsync(WebSocketCloseStatus.NormalClosure, "Выход", CancellationToken.None);
					break;
				}

				await SendMessageAsync(message);
			}

			await receiveTask;
		}

		private static async Task ConnectToServerAsync(string uri)
		{
			try
			{
				await _webSocket.ConnectAsync(new Uri(uri), CancellationToken.None);
				Console.WriteLine("Подключение к серверу...");
			}
			catch (Exception ex)
			{
				Console.WriteLine($"Ошибка подключения: {ex.Message}");
			}
		}

		private static async Task SendMessageAsync(string message)
		{
			var messageBuffer = Encoding.UTF8.GetBytes(message);
			await _webSocket.SendAsync(new ArraySegment<byte>(messageBuffer), WebSocketMessageType.Text, true, CancellationToken.None);
			Console.WriteLine($"Отправлено: {message}");
		}

		private static async Task ReceiveMessagesAsync()
		{
			var buffer = new byte[1024 * 4];

			try
			{
				while (_webSocket.State == WebSocketState.Open)
				{
					var result = await _webSocket.ReceiveAsync(new ArraySegment<byte>(buffer), CancellationToken.None);

					if (result.MessageType == WebSocketMessageType.Close)
					{
						await _webSocket.CloseAsync(WebSocketCloseStatus.NormalClosure, "Сервер закрыл соединение", CancellationToken.None);
						Console.WriteLine("Сервер закрыл соединение.");
						break;
					}

					var message = Encoding.UTF8.GetString(buffer, 0, result.Count);
					Console.WriteLine($"Получено: {message}");
				}
			}
			catch (WebSocketException ex)
			{
				Console.WriteLine($"Ошибка WebSocket: {ex.Message}");
			}
		}
	}
}
