# Refreshing filter function in Python

numbers = [i+1 for i in range(10)]
print("List of numbers", numbers)
print("Numbers greater than 5", list(filter(lambda x: x > 5, numbers)))