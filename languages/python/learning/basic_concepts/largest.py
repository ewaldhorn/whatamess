# .sort() sorts the list in-place
def get_largest(numbers,n):
  numbers.sort()
  return numbers[-n:]

# .sorted() returns a new, sorted copy of the list
def get_largest_safe(numbers,n):
  return sorted(numbers)[-n:]

sample = [2,3,4,1,34,123,321,1]

print("Sample:",sample)
result = get_largest(sample,3)
print("Result:",result)
print("Sample:",sample)

print("\n\n")

otherSample = [2,3,4,1,34,123,321,1]
print("Sample:",otherSample)
result = get_largest_safe(otherSample,3)
print("Result:",result)
print("Sample:",otherSample)
