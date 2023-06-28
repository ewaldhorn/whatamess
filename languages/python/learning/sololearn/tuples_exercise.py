contacts = [
    ('James', 42),
    ('Amy', 24),
    ('John', 31),
    ('Amanda', 63),
    ('Bob', 18)
]

data = input()
found = False

for c in contacts:
    if c[0] == data:
        print(c[0],"is",c[1])
        found = True

if found == False:
    print("Not Found")