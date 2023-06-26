using System;

public class Program
{
    static void Main(string[] args)
    {
        int num = Convert.ToInt32(Console.ReadLine());

        //your code goes here
        int sum = 0;

        while (num > 0) {
            sum += num;
            num--;
        }

        Console.WriteLine(sum);
        
    }
}
