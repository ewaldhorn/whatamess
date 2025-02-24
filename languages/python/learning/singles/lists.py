# create a list from a tuple
my_tuple = (1, 2, 3, 4, 5)
my_list = list(my_tuple)
print("Tuple->List", my_list)  # Output: [1, 2, 3, 4, 5]

# create a list from a set
my_set = {1, 2, 3, 4, 5}
my_list = list(my_set)
print("Set->List", my_list)  # Output: [1, 2, 3, 4, 5] (order may vary)

# dictionary to list
my_dict = {"a": 1, "b": 2, "c": 3}
my_list = list(my_dict)
print("Dictionary->List", my_list)  # Output: ['a', 'b', 'c'] (list of keys)
