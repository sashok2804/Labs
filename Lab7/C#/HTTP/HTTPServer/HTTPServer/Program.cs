using System;
using System.IO;
using System.Net;
using System.Text;
using System.Threading.Tasks;

class Program
{
	static async Task Main(string[] args)
	{
		HttpListener listener = new HttpListener();
		listener.Prefixes.Add("http://localhost:5000/"); 
		listener.Start();
		Console.WriteLine("Сервер запущен на http://localhost:5000/");

		while (true)
		{
			// Ожидаем входящий запрос
			HttpListenerContext context = await listener.GetContextAsync();

			// Логируем запрос
			LogRequest(context);

			// Маршрутизация
			if (context.Request.HttpMethod == "GET" && context.Request.Url.AbsolutePath == "/hello")
			{
				await HandleHelloRequest(context);
			}
			else if (context.Request.HttpMethod == "POST" && context.Request.Url.AbsolutePath == "/data")
			{
				await HandleDataRequest(context);
			}
			else
			{
				// Обработка неизвестного маршрута
				context.Response.StatusCode = (int)HttpStatusCode.NotFound;

				context.Response.ContentType = "text/plain; charset=utf-8";
				byte[] errorMessage = Encoding.UTF8.GetBytes("Маршрут не найден");

				await context.Response.OutputStream.WriteAsync(errorMessage, 0, errorMessage.Length);
				context.Response.Close();
			}
		}
	}

	// Метод для логирования запросов
	static void LogRequest(HttpListenerContext context)
	{
		string method = context.Request.HttpMethod;
		string url = context.Request.Url.ToString();
		DateTime startTime = DateTime.UtcNow;

		DateTime endTime = DateTime.UtcNow;
		double duration = ( endTime - startTime ).TotalMilliseconds;

		Console.WriteLine($"Метод: {method}, URL: {url}, Время выполнения: {duration} ms");
	}

	// Обработка GET запроса на /hello
	static async Task HandleHelloRequest(HttpListenerContext context)
	{
		string responseString = "Привет, мир!";
		byte[] buffer = Encoding.UTF8.GetBytes(responseString);

		context.Response.ContentType = "text/plain; charset=utf-8";
		context.Response.ContentLength64 = buffer.Length;
		context.Response.StatusCode = (int)HttpStatusCode.OK;

		await context.Response.OutputStream.WriteAsync(buffer, 0, buffer.Length);
		context.Response.Close();
	}

	// Обработка POST запроса на /data
	static async Task HandleDataRequest(HttpListenerContext context)
	{
		using (var reader = new StreamReader(context.Request.InputStream, Encoding.UTF8))
		{
			string jsonData = await reader.ReadToEndAsync();
			Console.WriteLine($"Полученные данные: {jsonData}");
		}

		context.Response.StatusCode = (int)HttpStatusCode.OK;
		context.Response.Close();
	}
}
