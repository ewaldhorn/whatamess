using System;

public class Program
{
  static void Main(string[] args)
    {
      int choice = 2;

      switch(choice)
      {
        case 1:
          Console.WriteLine("Sports");
          break;
            
        case 2:
          //match!
          Console.WriteLine("Business");
          break;

        case 3:
          Console.WriteLine("Technology");
          break;            
      }
  }
}
