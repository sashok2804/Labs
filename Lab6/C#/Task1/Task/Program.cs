using System;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

class Program
{
	// Функция для расчёта факториала числа
	static async Task<int> CalculateFactorialAsync(int n)
	{
		Console.WriteLine($"Расчёт факториала для числа {n}...");
		await Task.Delay(2000); // Имитация задержки выполнения
		int factorial = 1;
		for (int i = 2; i <= n; i++)
		{
			factorial *= i; // Вычисление факториала
		}
		return factorial;
	}

	// Функция для генерации случайных чисел
	static async Task<int[]> GenerateRandomNumbersAsync(int count)
	{
		Console.WriteLine("Генерация случайных чисел...");
		await Task.Delay(1000); // Имитация задержки выполнения
		Random rand = new Random();
		int[] randomNumbers = new int[count];
		for (int i = 0; i < count; i++)
		{
			randomNumbers[i] = rand.Next(0, 100); // Генерация случайного числа от 0 до 99
		}
		return randomNumbers;
	}

	// Функция для вычисления суммы числового ряда
	static async Task<int> CalculateSumAsync(int n)
	{
		Console.WriteLine($"Вычисление суммы числового ряда от 1 до {n}...");
		await Task.Delay(3000); // Имитация задержки выполнения
		int sum = Enumerable.Range(1, n).Sum(); // Вычисление суммы ряда
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
