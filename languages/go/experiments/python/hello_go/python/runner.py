import ctypes

lib = ctypes.CDLL("../bin/thelib.so")

# Important to set the correct return type, otherwise we get wrong answers
lib.CalculateLargeNumber.restype = ctypes.c_longlong

# header
print()
msg = "Python calling into a Go shared library."
print(msg)
print("-"*len(msg))

# content
bigNumber = lib.CalculateLargeNumber()
print("Go says the big number is      :", bigNumber)

# footer
print("And this is Python saying it is:",f"{bigNumber:,}".replace(",", " "))
print()
