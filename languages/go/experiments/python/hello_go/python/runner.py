import ctypes

lib = ctypes.CDLL("../bin/thelib.so")

# Important to set the correct return type, otherwise we get wrong answers
lib.CalculateLargeNumber.restype = ctypes.c_longlong

print()
bigNumber = lib.CalculateLargeNumber()
print("Go says the big number is      :", bigNumber)

print("And this is Python saying it is:",f"{bigNumber:,}".replace(",", " "))
print()
