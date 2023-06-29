def my_min(x, *y):
    smallest = x
    for i in y:
        if i < smallest:
            smallest = i
    return smallest

print(my_min(8, 13, 4, 42, 120, 7))