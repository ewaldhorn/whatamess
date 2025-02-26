# Just refreshing my list comprehension in Python.

# create a list of numbers from 1..10, i+1 to account for ranges being 0-based
numbers = [i+1 for i in range(10)]
print("List of numbers", numbers)

# create a list of lists using list comprehension
lists = [[j+1 for j in range(3)] for i in range(5)]
print("List of lists", lists)
