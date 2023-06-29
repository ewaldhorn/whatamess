myfile = open("filename.txt")

# write mode
open("filename.txt", "w")

# read mode
open("filename.txt", "r")
open("filename.txt")

# binary write mode
open("filename.txt", "wb")

file = open("filename.txt", "w")
# do stuff to the file
file.close()

file = open("/usercode/files/books.txt")
cont = file.read()
print(cont)
file.close()

# read X number of bytes per operation
file = open("/usercode/files/books.txt")
print(file.read(5))
print(file.read(7))
print(file.read())
file.close()

# get input as N, use that to determine how many characters to read
file = open("/usercode/files/books.txt")
#your code goes here
N = int(input())

print(file.read(N))

file.close()