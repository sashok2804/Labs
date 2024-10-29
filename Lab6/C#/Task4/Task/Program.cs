using System;
using System.Threading;
using System.Threading.Tasks;

class Program
{
	static int counter = 0;
	static Mutex mutex = new Mutex(); 

	static async Task Main(string[] args)
	{
		int numTasks = 5; 
		int numIterations = 1000; 

		// Создаём и запускаем задачи
		Task[] tasks = new Task[numTasks];
		for (int i = 0; i < numTasks; i++)
		{
			tasks[i] = Task.Run(() => IncrementCounter(numIterations));
		}

		await Task.WhenAll(tasks);

		Console.WriteLine($"Итоговое значение счётчика: {counter}");
	}

	static void IncrementCounter(int iterations)
	{
		for (int i = 0; i < iterations; i++)
		{
			//mutex.WaitOne();
			counter++;
			//mutex.ReleaseMutex();
		}
	}
}
