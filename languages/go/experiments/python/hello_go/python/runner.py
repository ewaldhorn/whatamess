import ctypes

lib = ctypes.CDLL("../bin/thelib.so")
lib.HelloFromGo()

print("And this is Python!")
