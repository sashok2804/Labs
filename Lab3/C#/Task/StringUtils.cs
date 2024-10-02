namespace StringUtils
{
	public static class StringHelper
	{
		public static string ReverseString(string input)
		{
			char[] charArray = input.ToCharArray();
			Array.Reverse(charArray);
			return new string(charArray);
		}
	}
}
