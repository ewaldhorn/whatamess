names = ["Bob", "Melanie", "Grace", "Tim", "Dan", "Leila", "Karen"]

def say_hello(who: str):
    print(f'Hello, {who}')

for i in range(len(names)):
    say_hello(names[i])

