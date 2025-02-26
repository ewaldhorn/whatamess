# Exploring map function in Python

numbers = [i+1 for i in range(10)]
print("List of numbers          :", numbers)
print("List of numbers (doubled):", list(map(lambda x: x*2, numbers)))