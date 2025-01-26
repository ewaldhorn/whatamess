import ctypes
import unittest

lib = ctypes.CDLL("../bin/thelib.so")

# Important to set the correct return type, otherwise we get wrong answers
lib.CalculateLargeNumber.restype = ctypes.c_longlong


class TestCalculations(unittest.TestCase):
  def test_calculateLargeNumber(self):
    result = lib.CalculateLargeNumber()
    self.assertEqual(result, 15432092901235)

if __name__ == '__main__':
  unittest.main()
