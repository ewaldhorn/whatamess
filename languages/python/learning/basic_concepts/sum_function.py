# Refresh Python sum function

numbers = [1,4,5,2,3]
print(sum(numbers))
print(sum(numbers,-8)) # give a starting point, default is 0

sorted_numbers = sorted(numbers)
print("The numbers are:", numbers)
print("Sorted are     :", sorted_numbers)
print("Descending are :", sorted(numbers,reverse=True))

contestants = [{"name":"Joe", "score":11},{"name":"Sandra","score":9},{"name":"Bob","score":12},{"name":"Alice","score":8}]
print("Contestants",contestants)
print("Sorted by name",sorted(contestants,key=lambda x: x["name"]))
print("Sorted by score",sorted(contestants,key=lambda x: x["score"],reverse=True))