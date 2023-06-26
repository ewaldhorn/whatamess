using System;

public class Program
{
    static void Main(string[] args)
    {
        Welcome("John");
        Welcome("Tom");
    }
    static void Welcome(string name) 
    {
        Console.WriteLine("Welcome " + name);
    }
}
