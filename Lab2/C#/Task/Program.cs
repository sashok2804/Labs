using System;
using System.Globalization;

class Program
{
	static void Main(string[] args)
	{
		// Создаем флаг task с параметром по умолчанию 1
		int task = 1;
		for (int i = 0; i < args.Length; i++)
		{
			if (args[i] == "-t" && i + 1 < args.Length)
			{
				task = int.Parse(args[i + 1]); // Получаем значение после флага -t
				break;
			}
		}

		// Выбор задания на основе значения флага
		switch (task)
		{
			case 1:
			Task1();
			break;
			case 2:
			Task2();
			break;
			case 4:
			Task4();
			break;
			case 5:
			Task5();
			break;
			case 6:
			Task6();
			break;
			default:
			Task1();
			break;
		}
	}

	static void Task1()
	{
		// Получение текущей даты и времени
		DateTime now = DateTime.Now;

		// Вывод даты
		Console.WriteLine($"Today {now.Day} {now.ToString("MMMM", CultureInfo.InvariantCulture)}, {now.Year} year.");

		// Вывод времени
		Console.WriteLine($"Current time: {now.Hour}:{now.Minute}:{now.Second}.");
	}

	static void Task2()
	{
		int i = 100;            // целое число
		float f = 3.14f;        // дробное число
		double d = 3.2314;      // дробное число с большей точностью
		bool b = true;          // булево значение
		string s = "Hello";     // строка

		// Вывод типов и значений
		Console.WriteLine($"{i.GetType()} - {i}");
		Console.WriteLine($"{f.GetType()} - {f}");
		Console.WriteLine($"{d.GetType()} - {d:F2}"); // Округление до 2 знаков
		Console.WriteLine($"{b.GetType()} - {b}");
		Console.WriteLine($"{s.GetType()} - {s}");
	}

	static void Task4()
	{
		int y = 33;
		int x = 44;

		// Вывод суммы
		Console.WriteLine($"X: {x}, Y: {y}, sum: {x + y}");
	}

	static void Task5()
	{
		double x = 124.17684;
		double y = 432.48724;

		// Вывод суммы и разности
		Console.WriteLine($"X: {x}, Y: {y}\nSum: {Plus(x, y)}, dif: {Minus(x, y)}");
	}

	static void Task6()
	{
		int[] arr = { 1, 2, 97 }; // объявляем массив
		int sum = 0;

		// Вывод массива
		Console.Write("Array: ");
		foreach (var num in arr)
		{
			Console.Write(num + " ");
			sum += num;
		}

		// Вычисление и вывод среднего значения
		double avg = (double)sum / arr.Length;
		Console.WriteLine($"\nAVG: {avg:F2}");
	}

	// Функция сложения двух чисел
	static double Plus(double a, double b)
	{
		return a + b;
	}

	// Функция вычитания двух чисел
	static double Minus(double a, double b)
	{
		return a - b;
	}
}
