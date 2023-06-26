using System;

public class Program
{
    static void Main(string[] args)
    {
        int num = 1;

        for(int i = 1; i<=100; i++)
        {
            //your code goes here
            if (num > 10000) {
                break;
            }


            num*=i;
        }
        Console.WriteLine(num);
    }
}
