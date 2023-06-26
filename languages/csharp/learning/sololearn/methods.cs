using System;

public class Program
{
    static void Main(string[] args)
    {
        //call the method
        Welcome(Console.ReadLine());
        
    }

    static void Welcome(string name)
    {
        //redesign the method
        Console.WriteLine($"Welcome, {name}!");
    }
}
