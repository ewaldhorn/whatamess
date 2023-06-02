(import (rnrs))

(define (middleOf start end)
  (div (+ start end) 2)
)

(define (searcher array target start end)
  (if (> start end)
    'not-found    
    (let (
        (midpoint (middleOf start end))
        (value (vector-ref array (middleOf start end)))
      )        
      (cond 
        ((= value target) midpoint)
        ((> value target) (searcher array target start (- midpoint 1)))
        ((< value target) (searcher array target (+ midpoint 1) end))
      )
    )   
  )
)

(define (binary-search array target)
   (searcher array target 0 (- (vector-length array) 1))
)
