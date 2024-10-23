using System;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

class Program
{
	static async Task<int> CalculateFactorialAsync(int n)
	{
		Console.WriteLine($"Расчёт факториала для числа {n}...");
		await Task.Delay(2000); // Имитация задержки выполнения
		int factorial = 1;
		for (int i = 2; i <= n; i++)
		{
			factorial *= i; 
		}
		return factorial;
	}

	static async Task<int[]> GenerateRandomNumbersAsync(int count)
	{
		Console.WriteLine("Генерация случайных чисел...");
		await Task.Delay(1000); 
		Random rand = new Random();
		int[] randomNumbers = new int[count];
		for (int i = 0; i < count; i++)
		{
			randomNumbers[i] = rand.Next(0, 100); 
		}
		return randomNumbers;
	}

	static async Task<int> CalculateSumAsync(int n)
	{
		Console.WriteLine($"Вычисление суммы числового ряда от 1 до {n}...");
		await Task.Delay(3000); 
		int sum = Enumerable.Range(1, n).Sum(); 
		return sum;
	}

	static async Task Main(string[] args)
	{
		// Запуск всех задач параллельно
		Task<int> factorialTask = CalculateFactorialAsync(5);
		Task<int[]> randomNumbersTask = GenerateRandomNumbersAsync(5);
		Task<int> sumTask = CalculateSumAsync(10);

		// Массив всех задач
		Task[] tasks = { factorialTask, randomNumbersTask, sumTask };

		// Пока есть незавершённые задачи
		while (tasks.Length > 0)
		{
			// Ожидаем первую завершившуюся задачу
			Task finishedTask = await Task.WhenAny(tasks);

			if (finishedTask == factorialTask)
			{
				int factorialResult = await factorialTask;
				Console.WriteLine($"Факториал: {factorialResult}");
			}
			else if (finishedTask == randomNumbersTask)
			{
				int[] randomNumbers = await randomNumbersTask;
				Console.WriteLine($"Случайные числа: {string.Join(", ", randomNumbers)}");
			}
			else if (finishedTask == sumTask)
			{
				int sumResult = await sumTask;
				Console.WriteLine($"Сумма числового ряда: {sumResult}");
			}

			// Обновляем массив, исключая завершённую задачу
			tasks = tasks.Where(t => !t.IsCompleted).ToArray();
		}
	}
}
