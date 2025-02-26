# Exploring the print function in Python beyond the basic usage.

numbers = [i+1 for i in range(10)]
print("List of numbers", numbers)
print("List of numbers", numbers, sep=" => ")
print("List of numbers", numbers, sep=" => ", end="\nDone\n")
