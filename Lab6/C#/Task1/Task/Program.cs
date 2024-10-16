using System;
using System.Collections.Generic;
using System.Drawing;

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
			default:
			Task1();
			break;
		}
	}
	static void Task1()
	{
		Person p = new Person("Alice", 30);
		p.Info();

		p.Birthday();
		p.Info();
	}

	static void Task2()
	{
		Circle c = new Circle(5);
		Console.WriteLine($"Circle Area: {c.Area():F2}");
	}

	static void Task3()
	{
		IShape c = new Circle(5);
		IShape r = new Rectangle(4, 5);

		List<IShape> shapes = new List<IShape> { c, r };

		foreach (var shape in shapes)
		{
			Console.WriteLine($"Area: {shape.Area():F2}");
		}
	}

	static void Task4()
	{
		Book b = new Book("C# top", "yandex");
		Console.WriteLine(b.String());
	}

}
public struct Person
{
	public string Name { get; }
	public int Age { get; private set; }

	public Person(string name, int age)
	{
		Name = name;
		Age = age;
	}

	public void Info()
	{
		Console.WriteLine($"Name: {Name}, Age: {Age}");
	}

	public void Birthday()
	{
		Age++;
	}
}

public struct Circle : IShape
{
	public double Radius { get; }

	public Circle(double radius)
	{
		Radius = radius;
	}

	public double Area()
	{
		return Math.PI * Radius * Radius;
	}
}

public struct Rectangle : IShape
{
	public double Width { get; }
	public double Height { get; }

	public Rectangle(double width, double height)
	{
		Width = width;
		Height = height;
	}

	public double Area()
	{
		return Width * Height;
	}
}

public interface IShape
{
	double Area();
}

public struct Book : IStringer
{
	public string Title { get; }
	public string Author { get; }

	public Book(string title, string author)
	{
		Title = title;
		Author = author;
	}

	public string String()
	{
		return $"Title: {Title}, Author: {Author}";
	}
}

public interface IStringer
{
	string String();
}
