import ctypes

lib = ctypes.CDLL("../bin/thelib.so")

print()
lib.HelloFromGo()
print(lib.CalculateLargeNumber())
print("Go says the big number is:", f"{lib.CalculateLargeNumber():,}".replace(",", " "))

print("And this is Python!")
print()
