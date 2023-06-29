file = open("newfile.txt", "w")
file.write("This has been written to a file")
file.close()


file = open("/usercode/files/books.txt", "a")

file.write("\nThe Da Vinci Code")
file.close()

file = open("/usercode/files/books.txt", "r")
print(file.read())
file.close()


## 
n = int(input())

file = open("numbers.txt", "w+")
#your code goes here
for i in range(1,n+1):
    file.write(str(i))
    file.write("\n")
file.close()


f = open("numbers.txt", "r")
print(f.read())
f.close()