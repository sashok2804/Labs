using System;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;

class TcpClientProgram
{
	public static async Task Main(string[] args)
	{
		using TcpClient client = new TcpClient();

		try
		{
			await client.ConnectAsync("127.0.0.1", 5000); 
			Console.WriteLine("Подключено к серверу.");

			using var stream = client.GetStream();

			// Отправка сообщения
			Console.Write("Введите сообщение для отправки серверу: ");
			string message = Console.ReadLine();
			byte[] data = Encoding.UTF8.GetBytes(message);
			await stream.WriteAsync(data, 0, data.Length);
			Console.WriteLine("Сообщение отправлено серверу.");

			// Чтение ответа сервера
			byte[] buffer = new byte[1024];
			Task readTask = Task.Run(async () =>
			{
				while (client.Connected)
				{
					try
					{
						if (stream.DataAvailable)
						{
							int bytesRead = await stream.ReadAsync(buffer, 0, buffer.Length);
							if (bytesRead == 0)
							{
								Console.WriteLine("Сервер закрыл соединение.");
								break; 
							}

							string response = Encoding.UTF8.GetString(buffer, 0, bytesRead);
							Console.WriteLine($"Ответ от сервера: {response}");

							// Проверяем сообщение о закрытии сервера
							if (response.Contains("Сервер закрывается"))
							{
								Console.WriteLine("Получено уведомление о завершении работы сервера.");
								break;
							}
						}

						await Task.Delay(100); // Задержка для проверки соединения
					}
					catch (Exception ex)
					{
						Console.WriteLine($"Ошибка при чтении данных от сервера: {ex.Message}");
						break;
					}
				}
			});

			await readTask;
		}
		catch (SocketException)
		{
			Console.WriteLine("Не удалось подключиться к серверу.");
		}
		catch (Exception ex)
		{
			Console.WriteLine($"Ошибка клиента: {ex.Message}");
		}
		finally
		{
			Console.WriteLine("Соединение закрыто.");
		}
	}
}
