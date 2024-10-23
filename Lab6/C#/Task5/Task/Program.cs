using System;
using System.Collections.Concurrent;
using System.Threading.Tasks;

class Program
{
	// Описание структуры запроса для калькулятора
	public class CalcRequest
	{
		public double Operand1 { get; set; }
		public double Operand2 { get; set; }
		public string Operation { get; set; }
		public TaskCompletionSource<double> Result { get; set; }
		public TaskCompletionSource<string> Error { get; set; }
	}

	// Функция-калькулятор, обрабатывающая запросы
	static void Calculator(ConcurrentQueue<CalcRequest> requests)
	{
		while (requests.TryDequeue(out CalcRequest req))
		{
			try
			{
				double result = 0;

				// Обработка операций
				switch (req.Operation)
				{
					case "+":
					result = req.Operand1 + req.Operand2;
					break;
					case "-":
					result = req.Operand1 - req.Operand2;
					break;
					case "*":
					result = req.Operand1 * req.Operand2;
					break;
					case "/":
					if (req.Operand2 == 0)
					{
						throw new DivideByZeroException("Деление на ноль");
					}
					result = req.Operand1 / req.Operand2;
					break;
					default:
					throw new InvalidOperationException("Неизвестная операция");
				}

				// Если ошибок нет, устанавливаем результат
				req.Result.SetResult(result);
			}
			catch (Exception ex)
			{
				// Если произошла ошибка, передаем её
				req.Error.SetResult(ex.Message);
			}
		}
	}

	// Функция для отправки запроса и получения результата
	static async Task SendRequest(double operand1, double operand2, string operation, ConcurrentQueue<CalcRequest> requests)
	{
		var resultSource = new TaskCompletionSource<double>();
		var errorSource = new TaskCompletionSource<string>();

		var request = new CalcRequest
		{
			Operand1 = operand1,
			Operand2 = operand2,
			Operation = operation,
			Result = resultSource,
			Error = errorSource
		};

		// Добавляем запрос в очередь
		requests.Enqueue(request);

		// Запускаем калькулятор для обработки запроса
		Calculator(requests);

		// Ожидаем результат или ошибку
		var completedTask = await Task.WhenAny(resultSource.Task, errorSource.Task);

		if (completedTask == resultSource.Task)
		{
			Console.WriteLine($"Результат: {operand1} {operation} {operand2} = {await resultSource.Task}");
		}
		else
		{
			Console.WriteLine($"Ошибка: {await errorSource.Task}");
		}
	}

	static async Task Main(string[] args)
	{
		var requests = new ConcurrentQueue<CalcRequest>();

		// Отправка запросов
		await SendRequest(5, 3, "+", requests);
		await SendRequest(7, 2, "-", requests);
		await SendRequest(6, 3, "*", requests);
		await SendRequest(10, 0, "/", requests);
		await SendRequest(9, 3, "/", requests);

		Console.ReadLine(); // Ожидание ввода для завершения программы
	}
}
