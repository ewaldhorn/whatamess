# Exploring map function in Python

numbers = [i+1 for i in range(10)]
print("List of numbers          :", numbers)
print("List of numbers (doubled):", list(map(lambda x: x*2, numbers)))

def tripler(n):
    return n * 3

tripled_numbers = map(tripler, numbers)
print("List of numbers :", numbers)
print("Numbers tripled :", list(tripled_numbers))

