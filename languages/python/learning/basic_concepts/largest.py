def get_largest(numbers,n):
  numbers.sort()
  return numbers[-n:]

sample = [2,3,4,1,34,123,321,1]

print("Sample:",sample)
result = get_largest(sample,3)
print("Result:",result)
print("Sample:",sample)
