using System;

public class Program
{
    static void Main(string[] args)
    {
        int num = Convert.ToInt32(Console.ReadLine());

        //your code goes here
        int answer = 1;

        for (int i = 2; i <= num; i++) {
            answer *= i;
        }

        Console.WriteLine(answer);
        
    }
}
