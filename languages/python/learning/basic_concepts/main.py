# Immutable
# str
# int
# float
# bool
# bytes
# tuple

x = (1, 2)
# x[0] = 1 # won't work, can't assign to tuple
y = x
x = (1, 2, 3)
print("Immutables:", x, y)

# Mutable
# list
# set
# dictionary
a = [1, 2]
b = a
a[0] = 4
print("Mutables:", a, b)
