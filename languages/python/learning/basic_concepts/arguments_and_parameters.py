def some_function(left, right, swap = False):
    if swap:
        left, right = right, left   

    print("Left:",left,"Right:",right)


def sum_them_all(*nums):
    total = 0
    for num in nums:
        total += num
    return total   

some_function(4,5)
some_function(right=5,left=4)
some_function(right=5,left=4, swap=True)
some_function(10,11)
some_function(10,11,True)
some_function(*[7,8]) # use * to destructure array

print("\n")
nums = [1]
print("The sum of",nums,"is",sum_them_all(*nums))
nums = [1,1]
print("The sum of",nums,"is",sum_them_all(*nums))
nums = [1,2]
print("The sum of",nums,"is",sum_them_all(*nums))
nums = [1,2,3]
print("The sum of",nums,"is",sum_them_all(*nums))
nums = [1,1,2,2]
print("The sum of",nums,"is",sum_them_all(*nums))
nums = [1,2,1,2]
print("The sum of",nums,"is",sum_them_all(*nums))


def multi_keyword_args(**kwargs):
    for key, value in kwargs.items():
        print(key, value)

print("\n")
multi_keyword_args()
multi_keyword_args(name="John", age=30)
multi_keyword_args(name="John", age=30, city="New York")    
