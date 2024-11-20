using System;

class Program
{
	static void Main(string[] args)
	{
		// Обрабатываем флаг -t, по умолчанию значение 1
		int task = 1;
		for (int i = 0; i < args.Length; i++)
		{
			if (args[i] == "-t" && i + 1 < args.Length)
			{
				task = int.Parse(args[i + 1]); // Получаем значение после флага -t
				break;
			}
		}

		// Вызываем задачу в зависимости от переданного значения
		switch (task)
		{
			case 1:
			Task1();
			break;
			case 2:
			Task2();
			break;
			case 3:
			Task3();
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
		Console.Write("Введите ваше число: ");
		int age = int.Parse(Console.ReadLine());

		if (age % 2 == 0)
		{
			Console.WriteLine($"{age} - четное");
		}
		else
		{
			Console.WriteLine($"{age} - нечетное");
		}
	}

	static void Task2()
	{
		int num = 0;
		Console.WriteLine($"{num} - {NumPos(num)}");
	}

	static void Task3()
	{
		int count = 11;
		for (int i = 1; i < count; i++)
		{
			Console.Write($"{i}, ");
		}
		Console.WriteLine();
	}

	static void Task4()
	{
		Console.Write("Введите строку: ");
		string alena = Console.ReadLine();
		Console.WriteLine($"Строка '{alena}' - {StrLen(alena)} симв.");
	}

	static void Task5()
	{
		Rectangle rect = new Rectangle(10, 20);
		rect.Square();
	}

	static void Task6()
	{
		int num1 = 30;
		int num2 = 15;
		Console.WriteLine($"first - {num1}, second - {num2}, avg - {Sred(num1, num2):F2}");
	}

	static string NumPos(int n)
	{
		if (n > 0)
			return "positive";
		if (n < 0)
			return "negative";
		return "zero";
	}

	static int StrLen(string s)
	{
		return s.Length;
	}

	static double Sred(int first, int second)
	{
		return ( first + second ) / 2.0;
	}
}

class Rectangle
{
	public int Width { get; set; }
	public int Height { get; set; }

	public Rectangle(int width, int height)
	{
		Width = width;
		Height = height;
	}

	public int Square()
	{
		int s = Width * Height;
		Console.WriteLine($"Ширина - {Width}, высота - {Height}, площадь - {s}");
		return s;
	}
}
