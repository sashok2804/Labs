namespace MathUtils
{
	public static class MathHelper
	{
		public static int Factorial(int n)
		{
			if (n < 0)
			{
				throw new ArgumentException("Число должно быть неотрицательным.");
			}

			int result = 1;
			for (int i = 1; i <= n; i++)
			{
				result *= i;
			}

			return result;
		}
	}
}
