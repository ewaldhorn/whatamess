proc isLeapYear*(year: int): bool =
  # must be divisible by 4, unless divisible by 100 AND 400
  if year mod 100 == 0:
    year mod 400 == 0
  else: year mod 4 == 0

