# Python does support type hinting
# but.... it's not enforced
number: int = 10
notANumber: str = "11"

number = "Sssss"
number = {3, 4, 5, 6, 7, 4, 5, 6, 7}

print(number)
print(notANumber)

from typing import Final

VERSION: Final[str] = "10.0.1"
print(VERSION)
