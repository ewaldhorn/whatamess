import ctypes

lib = ctypes.CDLL("../bin/thelib.so")

# Important to set the correct return type, otherwise we get wrong answers
lib.CalculateLargeNumber.restype = ctypes.c_longlong

print()
print(lib.CalculateLargeNumber())
print("Go says the big number is:", f"{lib.CalculateLargeNumber():,}".replace(",", " "))

print("And this is Python!")
print()
