using System;
using System.Collections.Concurrent;
using System.IO;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

class Program
{
	// Функция для реверсирования строки
	static string ReverseString(string s)
	{
		return new string(s.Reverse().ToArray());
	}

	static async Task Main()
	{
		// Открываем файл с данными
		var lines = await File.ReadAllLinesAsync("input.txt");
		var tasks = new ConcurrentBag<string>(lines);

		Console.Write("Enter number of workers: ");
		int workerCount = int.Parse(Console.ReadLine());

		var result = new ConcurrentBag<string>();
		var tasksList = new Task[workerCount];

		// Запускаем воркеров
		for (int i = 0; i < workerCount; i++)
		{
			tasksList[i] = Task.Run(() =>
			{
				while (tasks.TryTake(out var task))
				{
					Console.WriteLine($"Worker {Thread.CurrentThread.ManagedThreadId} processing task: {task}");
					result.Add(ReverseString(task));
				}
			});
		}

		await Task.WhenAll(tasksList);

		// Выводим результаты работы воркеров
		Console.WriteLine("Reversed lines:");
		foreach (var reversed in result)
		{
			Console.WriteLine(reversed);
		}
	}
}
