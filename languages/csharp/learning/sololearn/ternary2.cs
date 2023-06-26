using System;

public class Program
{
    static void Main(string[] args)
    {
        int hour = Convert.ToInt32(Console.ReadLine());

        //your code goes here
        Console.WriteLine(hour<=12?"AM":"PM");
        
    }
}
