using System;
using System.Collections.Concurrent;
using System.Threading;
using System.Threading.Tasks;

class Program
{
	static async Task Main(string[] args)
	{
		int n = 10;
		var channel = new BlockingCollection<int>();

		// Запускаем задачу для генерации чисел Фибоначчи
		var producerTask = Task.Run(() => GenerateFibonacci(n, channel));

		// Запускаем задачу для чтения и вывода чисел
		var consumerTask = Task.Run(() => PrintFibonacci(channel));

		// Ожидаем завершения задач
		await Task.WhenAll(producerTask, consumerTask);
	}

	static void GenerateFibonacci(int n, BlockingCollection<int> channel)
	{
		int a = 0, b = 1;
		for (int i = 0; i < n; i++)
		{
			channel.Add(a);
			int temp = a;
			a = b;
			b = temp + b;
		}
		channel.CompleteAdding(); // Сигнализируем, что данные больше не будут отправляться
	}

	static void PrintFibonacci(BlockingCollection<int> channel)
	{
		foreach (var num in channel.GetConsumingEnumerable())
		{
			Console.WriteLine(num);
		}
	}
}
