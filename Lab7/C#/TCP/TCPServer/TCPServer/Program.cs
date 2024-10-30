using System;
using System.Collections.Concurrent;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

class TcpServer
{
	private static TcpListener _listener;
	private static CancellationTokenSource _cancellationTokenSource = new CancellationTokenSource();
	private static ConcurrentBag<TcpClient> _clients = new ConcurrentBag<TcpClient>();

	public static async Task StartAsync(int port)
	{
		_listener = new TcpListener(IPAddress.Any, port);
		_listener.Start();
		Console.WriteLine($"Сервер запущен на порту {port}");

		try
		{
			while (!_cancellationTokenSource.Token.IsCancellationRequested)
			{
				var client = await _listener.AcceptTcpClientAsync();
				_clients.Add(client); 
				_ = Task.Run(() => HandleClientAsync(client, _cancellationTokenSource.Token));
			}
		}
		catch (Exception ex) when (_cancellationTokenSource.Token.IsCancellationRequested)
		{
			Console.WriteLine("Сервер остановлен.");
		}
		catch (Exception ex)
		{
			Console.WriteLine($"Ошибка сервера: {ex.Message}");
		}
	}

	private static async Task HandleClientAsync(TcpClient client, CancellationToken token)
	{
		Console.WriteLine("Клиент подключен");
		using var stream = client.GetStream();
		byte[] buffer = new byte[1024];

		try
		{
			while (!token.IsCancellationRequested && client.Connected)
			{
				if (stream.DataAvailable)
				{
					int bytesRead = await stream.ReadAsync(buffer, 0, buffer.Length, token);
					if (bytesRead == 0) break; // Разрыв соединения клиентом

					string message = Encoding.UTF8.GetString(buffer, 0, bytesRead);
					Console.WriteLine($"Получено сообщение от клиента: {message}");

					string response = "Сообщение получено";
					byte[] responseData = Encoding.UTF8.GetBytes(response);
					await stream.WriteAsync(responseData, 0, responseData.Length, token);
					Console.WriteLine("Ответ отправлен клиенту.");
				}

				await Task.Delay(100); 
			}
		}
		catch (Exception ex)
		{
			Console.WriteLine($"Ошибка при обработке клиента: {ex.Message}");
		}
		finally
		{
			client.Close();
			Console.WriteLine("Клиент отключен");
		}
	}

	public static async Task StopAsync()
	{
		_cancellationTokenSource.Cancel();
		_listener.Stop();
		Console.WriteLine("Сервер остановлен. Завершаем все активные соединения...");

		// Завершение всех активных соединений
		foreach (var client in _clients)
		{
			if (client.Connected)
			{
				try
				{
					using var stream = client.GetStream();
					string shutdownMessage = "Сервер закрывается";
					byte[] shutdownData = Encoding.UTF8.GetBytes(shutdownMessage);
					await stream.WriteAsync(shutdownData, 0, shutdownData.Length);
				}
				catch (Exception ex)
				{
					Console.WriteLine($"Ошибка при отправке сообщения об остановке клиенту: {ex.Message}");
				}
				finally
				{
					client.Close(); 
				}
			}
		}
		Console.WriteLine("Все активные соединения корректно завершены.");
	}
}

class Program
{
	static async Task Main(string[] args)
	{
		int port = 5000; 
		var serverTask = TcpServer.StartAsync(port);

		Console.WriteLine("Нажмите любую клавишу для остановки сервера...");
		Console.ReadKey();

		await TcpServer.StopAsync();
		await serverTask;
	}
}
