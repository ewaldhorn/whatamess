using System;

public class Program
{
    static void Main(string[] args)
    {
        Welcome("John", 25);
        Welcome("Tom", 36);
    }
    static void Welcome(string name, int age) 
    {
        Console.WriteLine("Welcome " + name);
        Console.WriteLine("You age: " + age);
    }
}
