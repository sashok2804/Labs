using System;
using System.Threading;
using System.Threading.Tasks;

class Program
{
	static int counter = 0;
	static Mutex mutex = new Mutex(); // Мьютекс для синхронизации

	static async Task Main(string[] args)
	{
		int numTasks = 5; // Количество задач (аналог горутин)
		int numIterations = 1000; // Количество итераций для каждой задачи

		// Создаём и запускаем задачи
		Task[] tasks = new Task[numTasks];
		for (int i = 0; i < numTasks; i++)
		{
			tasks[i] = Task.Run(() => IncrementCounter(numIterations));
		}

		// Ожидаем завершения всех задач
		await Task.WhenAll(tasks);

		// Выводим итоговое значение счётчика
		Console.WriteLine($"Итоговое значение счётчика: {counter}");
	}

	static void IncrementCounter(int iterations)
	{
		for (int i = 0; i < iterations; i++)
		{
			// Закомментируйте/разкомментируйте мьютекс, чтобы увидеть разницу
			mutex.WaitOne(); // Вход в критическую секцию (захват мьютекса)
			counter++;
			mutex.ReleaseMutex(); // Выход из критической секции (освобождение мьютекса)
		}
	}
}
