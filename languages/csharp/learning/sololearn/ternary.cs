using System;

public class Program
{
    static void Main(string[] args)
    {
        int age = 42;
        string isAdult = age<18 ? "Too young": "Old enough";

        Console.WriteLine(isAdult);
    }
}
