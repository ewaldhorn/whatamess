(import (rnrs))

(define (hamming-distance strand-a strand-b)
    (unless (= (string-length strand-a)(string-length strand-b))
    (error 'hamming "two strings must be of the same length"))
    (let loop ((distance 0) (i 0))
    (if (= i (string-length strand-a))
        distance
        (loop (+ distance (calc (string-ref strand-a i) (string-ref strand-b i))) (+ i 1)))))

(define (calc L R)
  (if (eq? L R)
      0
      1))


