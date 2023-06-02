(import (rnrs))
(define (square n)
  (if (or (< n 1) (> n 64)) (error "Invalid square specified.")
      (expt 2 (- n 1))))
(define total
  (+ -1 (expt 2 (- 64 0))))

