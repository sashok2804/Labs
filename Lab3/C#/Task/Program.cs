using System;
using MathUtils;
using StringUtils;
using System.Collections.Generic;

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
			default:
			Task1();
			break;
		}
	}

	static void Task1()
	{
		Console.Write("Введите число для вычисления факториала: ");
		int number = int.Parse(Console.ReadLine());
		try
		{
			int result = MathHelper.Factorial(number);
			Console.WriteLine($"Факториал числа {number}: {result}");
		}
		catch (ArgumentException ex)
		{
			Console.WriteLine(ex.Message);
		}
	}

	static void Task2()
	{
		Console.Write("Введите строку для переворота: ");
		string input = Console.ReadLine();
		string reversed = StringHelper.ReverseString(input);
		Console.WriteLine($"Перевернутая строка: {reversed}");
	}

	static void Task3()
	{
		int[] array = new int[5];
		for (int i = 0; i < array.Length; i++)
		{
			Console.Write($"Введите элемент {i + 1}: ");
			array[i] = int.Parse(Console.ReadLine());
		}
		Console.WriteLine("Вы ввели следующие числа:");
		foreach (var num in array)
		{
			Console.Write(num + " ");
		}
	}

	static void Task4()
	{

		List<int> list = new List<int> { 1, 2, 3, 4, 5 };
		Console.WriteLine("Исходный список: " + string.Join(", ", list));

		list.Add(6);
		Console.WriteLine("После добавления 6: " + string.Join(", ", list));

		list.Remove(3);
		Console.WriteLine("После удаления 3: " + string.Join(", ", list));
	}

	static void Task5()
	{

		List<string> stringList = new List<string> { "apple", "banana", "watermelon", "grape", "strawberry" };
		Console.WriteLine("Исходный список строк: " + string.Join(", ", stringList));

		string longest = stringList[0];
		foreach (var str in stringList)
		{
			if (str.Length > longest.Length)
			{
				longest = str;
			}
		}

		Console.WriteLine($"Самая длинная строка: {longest}");
	}
}
