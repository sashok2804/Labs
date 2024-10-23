using System;
using System.Collections.Concurrent;
using System.Threading;
using System.Threading.Tasks;

class Program
{
	static async Task Main(string[] args)
	{
		var numCh = new BlockingCollection<int>();
		var resultCh = new BlockingCollection<string>();

		// Запускаем генерацию случайных чисел
		var generateTask = Task.Run(() => GenerateNumbers(numCh));

		// Запускаем проверку на четность/нечетность
		var checkTask = Task.Run(() => CheckEvenOdd(numCh, resultCh));

		// Обработка сообщений из каналов
		var resultTask = Task.Run(() => ProcessResults(resultCh));

		// Ожидаем завершения
		await Task.WhenAll(generateTask, checkTask, resultTask);
	}

	static void GenerateNumbers(BlockingCollection<int> numCh)
	{
		var rand = new Random();
		while (true)
		{
			int num = rand.Next(100);
			numCh.Add(num);
			Thread.Sleep(1000); // Пауза для имитации задержки
		}
	}

	static void CheckEvenOdd(BlockingCollection<int> numCh, BlockingCollection<string> resultCh)
	{
		foreach (var num in numCh.GetConsumingEnumerable())
		{
			if (num % 2 == 0)
			{
				resultCh.Add($"Число {num} чётное");
			}
			else
			{
				resultCh.Add($"Число {num} нечётное");
			}
		}
	}

	static void ProcessResults(BlockingCollection<string> resultCh)
	{
		foreach (var result in resultCh.GetConsumingEnumerable())
		{
			Console.WriteLine(result);
		}
	}
}
