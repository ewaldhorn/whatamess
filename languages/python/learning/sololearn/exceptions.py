try:
    num1 = 7
    num2 = 0
    print (num1 / num2)
    print("Done calculation")
except ZeroDivisionError:
    print("An error occurred")
    print("due to zero division")


def withdraw(amount):
   print(str(amount) + " withdrawn!")

#your code goes here
try:
   amt = input()
   amount = int(amt)
   withdraw(amount)
except ValueError:
   print("Please enter a number")