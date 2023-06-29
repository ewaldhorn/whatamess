class Shape: 
    def __init__(self, w, h):
        self.width = w
        self.height = h

    def area(self):
        print(self.width*self.height)

class Rectangle(Shape):
    #your code goes here
    def perimeter(self):
        print((self.width+self.height)*2)

w = int(input())
h = int(input())

r = Rectangle(w, h)
r.area()
r.perimeter()