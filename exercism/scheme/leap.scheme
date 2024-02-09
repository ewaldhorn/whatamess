(import (rnrs))
(define (divisible? n x)
  (zero? (remainder n x)))
(define (leap-year? year)
    (if (and (divisible? year 4) (not (divisible? year 100)))
      #t
      (if (divisible? year 400)
          #t
          #f)))
