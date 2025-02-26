# Refreshing filter function in Python

numbers = [i+1 for i in range(10)]
print("List of numbers", numbers)
print("Numbers greater than 5", list(filter(lambda x: x > 5, numbers)))

def is_odd(b):
    return b % 2 == 1

print("Odd numbers", list(filter(is_odd, numbers)))