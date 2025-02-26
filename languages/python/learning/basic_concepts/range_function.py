# Exploring Python's range function a bit.

print("Just a range of numbers, from 0 to < 10")
nums = range(10)
print(nums)
print(list(nums))

print("\nSame thing, but starting at 1")
nums = range(1,10)
print(nums)
print(list(nums))

print("\nSame thing, but with a step of 2")
nums = range(1,10,2)
print(nums)
print(list(nums))