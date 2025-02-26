def some_function(left, right, swap = False):
    if swap:
        left, right = right, left   

    print("Left:",left,"Right:",right)

some_function(4,5)
some_function(right=5,left=4)
some_function(right=5,left=4, swap=True)
some_function(10,11)
some_function(10,11,True)