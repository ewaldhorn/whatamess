# Just refreshing my list comprehension in Python.

# create a list of numbers from 1..10, i+1 to account for ranges being 0-based
numbers = [i+1 for i in range(10)]
print("List of numbers", numbers)

# create a list of lists using list comprehension
lists = [[j+1 for j in range(3)] for i in range(5)]
print("List of lists", lists)

# create a list of numbers that are multiples of three
multiples_of_3 = [i for i in range(100) if i%3==0]
print("Multiples of 3", multiples_of_3)
