using System;
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
		var people = new Dictionary<string, int>();
		AddPerson(people, "Иван", 25);
		AddPerson(people, "Анна", 30);
		AddPerson(people, "Сергей", 35);

		Console.WriteLine("Список людей:");
		PrintPeople(people);
	}

	static void Task2()
	{
		var people = new Dictionary<string, int>();
		AddPerson(people, "Иван", 25);
		AddPerson(people, "Анна", 30);
		AddPerson(people, "Сергей", 35);

		Console.WriteLine($"Средний возраст: {AverageAge(people):F2}");
	}

	static void Task3()
	{
		var people = new Dictionary<string, int>();
		AddPerson(people, "Иван", 25);
		AddPerson(people, "Анна", 30);
		AddPerson(people, "Сергей", 35);

		Console.WriteLine("Список до удаления Анны:");
		PrintPeople(people);

		RemovePerson(people, "Анна");
		Console.WriteLine("Список после удаления Анны:");
		PrintPeople(people);
	}

	static void Task4()
	{
		Console.WriteLine("Введите строку:");
		string input = Console.ReadLine();
		Console.WriteLine(input.ToUpper());
	}

	static void Task5()
	{
		Console.WriteLine("Введите числа через пробел:");
		string input = Console.ReadLine();
		var numbers = input.Split(' ').Select(x => int.TryParse(x, out var num) ? num : 0).ToArray();
		int sum = numbers.Sum();
		Console.WriteLine($"Сумма чисел: {sum}");
	}

	static void Task6()
	{
		Console.WriteLine("Введите числа через пробел:");
		string input = Console.ReadLine();
		var numbers = input.Split(' ').Select(x => int.TryParse(x, out var num) ? num : 0).ToArray();

		Console.WriteLine("Массив в обратном порядке:");
		for (int i = numbers.Length - 1; i >= 0; i--)
		{
			Console.Write($"{numbers[i]} ");
		}
		Console.WriteLine();
	}

	static void AddPerson(Dictionary<string, int> people, string name, int age)
	{
		people[name] = age;
	}

	static void PrintPeople(Dictionary<string, int> people)
	{
		foreach (var person in people)
		{
			Console.WriteLine($"{person.Key}: {person.Value}");
		}
	}

	static void RemovePerson(Dictionary<string, int> people, string name)
	{
		people.Remove(name);
	}

	static double AverageAge(Dictionary<string, int> people)
	{
		if (people.Count == 0) return 0;

		return people.Values.Average();
	}
}
