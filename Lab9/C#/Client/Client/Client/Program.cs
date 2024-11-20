using System;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;
using System.Collections.Generic;

public class AuthRequest
{
	public string Username { get; set; }
	public string Password { get; set; }
}

public class AuthResponse
{
	public string Token { get; set; }
}

public class User
{
	public int Id { get; set; }
	public string Name { get; set; }
	public int Age { get; set; }
}

public class CreateUserRequest
{
	public int Id { get; set; }
	public string Name { get; set; }
	public int Age { get; set; }
}

class Program
{
	private static readonly HttpClient client = new HttpClient();
	private static string token = string.Empty;

	static async Task Main(string[] args)
	{
		bool isAuthenticated = false;

		while (true)
		{
			Console.WriteLine("Выберите опцию:");
			Console.WriteLine("1. Аутентификация");
			Console.WriteLine("2. Получить пользователей ");
			Console.WriteLine("3. Создать нового пользователя ");
			Console.WriteLine("4. Получить пользователя по ID ");
			Console.WriteLine("5. Обновить пользователя по ID ");
			Console.WriteLine("6. Удалить пользователя по ID ");
			Console.WriteLine("7. Выход");
			string choice = Console.ReadLine();

			switch (choice)
			{
				case "1":
				await Authenticate();
				isAuthenticated = true;
				break;

				case "2":
				if (isAuthenticated)
				{
					await GetUsers();
				}
				else
				{
					Console.WriteLine("Сначала необходимо пройти аутентификацию.");
				}
				break;

				case "3":
				if (isAuthenticated)
				{
					await CreateUser();
				}
				else
				{
					Console.WriteLine("Сначала необходимо пройти аутентификацию.");
				}
				break;

				case "4":
				if (isAuthenticated)
				{
					Console.WriteLine("Введите ID пользователя:");
					if (int.TryParse(Console.ReadLine(), out int userId))
					{
						await GetUserById(userId);
					}
					else
					{
						Console.WriteLine("Неверный ID пользователя.");
					}
				}
				else
				{
					Console.WriteLine("Сначала необходимо пройти аутентификацию.");
				}
				break;

				case "5":
				if (isAuthenticated)
				{
					Console.WriteLine("Введите ID пользователя для обновления:");
					if (int.TryParse(Console.ReadLine(), out int userId))
					{
						await UpdateUser(userId);
					}
					else
					{
						Console.WriteLine("Неверный ID пользователя.");
					}
				}
				else
				{
					Console.WriteLine("Сначала необходимо пройти аутентификацию.");
				}
				break;

				case "6":
				if (isAuthenticated)
				{
					Console.WriteLine("Введите ID пользователя для удаления:");
					if (int.TryParse(Console.ReadLine(), out int userId))
					{
						await DeleteUser(userId);
					}
					else
					{
						Console.WriteLine("Неверный ID пользователя.");
					}
				}
				else
				{
					Console.WriteLine("Сначала необходимо пройти аутентификацию.");
				}
				break;

				case "7":
				return;  

				default:
				Console.WriteLine("Неверный выбор, попробуйте снова.");
				break;
			}
		}
	}

	private static async Task Authenticate()
	{
		Console.Clear();

		Console.Write("Имя пользователя: ");
		string username = Console.ReadLine();

		Console.Write("Пароль: ");
		string password = ReadPassword();  

		// Создаем объект для аутентификации
		var authRequest = new AuthRequest
		{
			Username = username,
			Password = password
		};

		var jsonRequest = JsonConvert.SerializeObject(authRequest);
		var content = new StringContent(jsonRequest, Encoding.UTF8, "application/json");

		string url = "http://localhost:5222/api/Auth/login";

		try
		{
			HttpResponseMessage response = await client.PostAsync(url, content);

			// Проверка на успешный ответ
			if (response.IsSuccessStatusCode)
			{
				string responseBody = await response.Content.ReadAsStringAsync();
				var authResponse = JsonConvert.DeserializeObject<AuthResponse>(responseBody);
				token = authResponse.Token;
				Console.Clear();
				Console.WriteLine("Аутентификация прошла успешно! Токен получен.");
			}
			else
			{
				Console.Clear();
				Console.WriteLine($"Неудачная аутентификация: {response.StatusCode}, {response.ReasonPhrase}");
			}
		}
		catch (Exception ex)
		{
			Console.Clear();
			Console.WriteLine($"Ошибка при аутентификации: {ex.Message}");
		}
	}

	// Метод для скрытого ввода пароля
	private static string ReadPassword()
	{
		StringBuilder password = new StringBuilder();
		ConsoleKeyInfo key;
		while (( key = Console.ReadKey(true) ).Key != ConsoleKey.Enter)
		{
			if (key.Key == ConsoleKey.Backspace && password.Length > 0)
			{
				password.Length--; // Удаляем последний символ, если была нажата Backspace
				Console.Write("\b \b"); // Убираем символ на экране
			}
			else if (!char.IsControl(key.KeyChar))
			{
				password.Append(key.KeyChar); // Добавляем символ в пароль
				Console.Write("*"); // Показываем * вместо символов пароля
			}
		}
		Console.WriteLine();
		return password.ToString();
	}

	private static async Task GetUsers()
	{
		Console.Clear();

		string url = "http://localhost:5222/api/Users";

		try
		{
			// Убираем старое значение Authorization, если оно есть
			client.DefaultRequestHeaders.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", token);
			client.DefaultRequestHeaders.Add("accept", "text/plain");

			HttpResponseMessage response = await client.GetAsync(url);

			if (response.IsSuccessStatusCode)
			{
				string responseBody = await response.Content.ReadAsStringAsync();
				var users = JsonConvert.DeserializeObject<List<User>>(responseBody);

				Console.WriteLine("Список пользователей:\n");

				if (users.Count == 0)
				{
					Console.WriteLine("Нет пользователей в базе данных.");
				}
				else
				{
					foreach (var user in users)
					{
						Console.WriteLine($"Id: {user.Id}, Имя: {user.Name}, Возраст: {user.Age}");
					}
				}
			}
			else
			{
				Console.WriteLine($"\nНе удалось загрузить пользователей: {response.StatusCode}, {response.ReasonPhrase}");
			}
		}
		catch (Exception ex)
		{
			Console.WriteLine($"\nОшибка при загрузке пользователей: {ex.Message}");
		}
	}

	private static async Task CreateUser()
	{
		Console.Clear();

		Console.WriteLine("Введите данные для нового пользователя.\n");

		Console.Write("Имя: ");
		string name = Console.ReadLine();

		Console.Write("Возраст: ");
		int age;
		while (!int.TryParse(Console.ReadLine(), out age) || age <= 0)
		{
			Console.WriteLine("Введите корректный возраст (положительное число).");
			Console.Write("Возраст: ");
		}

		var createUserRequest = new CreateUserRequest
		{
			Id = 0,  // Сервер сгенерирует новый ID
			Name = name,
			Age = age
		};

		var jsonRequest = JsonConvert.SerializeObject(createUserRequest);
		var content = new StringContent(jsonRequest, Encoding.UTF8, "application/json");

		string url = "http://localhost:5222/api/Users";

		try
		{
			client.DefaultRequestHeaders.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", token);
			client.DefaultRequestHeaders.Add("accept", "text/plain");

			HttpResponseMessage response = await client.PostAsync(url, content);

			if (response.IsSuccessStatusCode)
			{
				string responseBody = await response.Content.ReadAsStringAsync();
				var createdUser = JsonConvert.DeserializeObject<User>(responseBody);

				Console.WriteLine("\nПользователь успешно создан!");
				Console.WriteLine($"Id: {createdUser.Id}, Имя: {createdUser.Name}, Возраст: {createdUser.Age}");
			}
			else
			{
				Console.WriteLine($"\nНе удалось создать пользователя: {response.StatusCode}, {response.ReasonPhrase}");
			}
		}
		catch (Exception ex)
		{
			Console.WriteLine($"\nОшибка при создании пользователя: {ex.Message}");
		}
	}

	private static async Task GetUserById(int userId)
	{
		Console.Clear();

		string url = $"http://localhost:5222/api/Users/{userId}";

		try
		{
			client.DefaultRequestHeaders.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", token);
			client.DefaultRequestHeaders.Add("accept", "text/plain");

			HttpResponseMessage response = await client.GetAsync(url);

			if (response.IsSuccessStatusCode)
			{
				string responseBody = await response.Content.ReadAsStringAsync();
				var user = JsonConvert.DeserializeObject<User>(responseBody);

				Console.WriteLine("\nДанные пользователя:");
				Console.WriteLine($"Id: {user.Id}, Имя: {user.Name}, Возраст: {user.Age}");
			}
			else
			{
				Console.WriteLine($"\nНе удалось получить пользователя с ID {userId}: {response.StatusCode}, {response.ReasonPhrase}");
			}
		}
		catch (Exception ex)
		{
			Console.WriteLine($"\nОшибка при получении пользователя: {ex.Message}");
		}
	}

	private static async Task UpdateUser(int userId)
	{
		string url = $"http://localhost:5222/api/Users/{userId}";

		Console.WriteLine("Введите новое имя пользователя:");
		string name = Console.ReadLine();  

		Console.WriteLine("Введите новый возраст пользователя:");
		int age;
		while (!int.TryParse(Console.ReadLine(), out age))
		{
			Console.WriteLine("Пожалуйста, введите корректный возраст.");
		}

		var updateUserRequest = new CreateUserRequest
		{
			Id = 0,  
			Name = name,
			Age = age
		};

		var jsonRequest = JsonConvert.SerializeObject(updateUserRequest);
		var content = new StringContent(jsonRequest, Encoding.UTF8, "application/json");

		try
		{
			client.DefaultRequestHeaders.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", token);
			client.DefaultRequestHeaders.Add("accept", "*/*");

			HttpResponseMessage response = await client.PutAsync(url, content);

			if (response.IsSuccessStatusCode)
			{
				Console.WriteLine($"Пользователь с ID {userId} успешно обновлен.");
			}
			else
			{
				Console.WriteLine($"Не удалось обновить пользователя с ID {userId}: {response.StatusCode}, {response.ReasonPhrase}");
			}
		}
		catch (Exception ex)
		{
			Console.WriteLine("Ошибка при обновлении пользователя: " + ex.Message);
		}
	}

	private static async Task DeleteUser(int userId)
	{
		string url = $"http://localhost:5222/api/Users/{userId}";

		try
		{
			client.DefaultRequestHeaders.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", token);
			client.DefaultRequestHeaders.Add("accept", "*/*");

			HttpResponseMessage response = await client.DeleteAsync(url);

			if (response.IsSuccessStatusCode)
			{
				Console.WriteLine($"Пользователь с ID {userId} успешно удален.");
			}
			else
			{
				Console.WriteLine($"Не удалось удалить пользователя с ID {userId}: {response.StatusCode}, {response.ReasonPhrase}");
			}
		}
		catch (Exception ex)
		{
			Console.WriteLine("Ошибка при удалении пользователя: " + ex.Message);
		}
	}
}
