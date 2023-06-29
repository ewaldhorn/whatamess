try:
  f = open("/usercode/files/books.txt")
  cont = f.read()
  print(cont)
finally:
 f.close()


# auto closes
with open("/usercode/files/books.txt") as f:
  print(f.read())

with open("/usercode/files/books.txt") as f:
   #your code goes here
   strings = f.read().split("\n")
   pos = 1
   for s in strings:
      count = len(s.split(" "))
      print(f"Line {pos}: {count} words")
      pos += 1

   