;; Exercism Darts
(module
  (func (export "score") (param $x f32) (param $y f32) (result i32)
    ;; create locals
    (local $xx f32)
    (local $yy f32)
    (local $radius f32)
    
    ;; calculate radius
    (local.set $xx (f32.mul (local.get $x)(local.get $x)))
    (local.set $yy (f32.mul (local.get $y)(local.get $y)))
    (local.set $radius (f32.add(local.get $xx)(local.get $yy)))

    ;; check the radius to see what the score should be
    ;; greater than 100 means score == 0
    local.get $radius
    f32.const 100.0
    f32.gt

    (if 
      (then (
        return (i32.const 0)
      )
    ))

    ;; greater than 25 means score == 1 
    local.get $radius
    f32.const 25.0
    f32.gt

    (if 
      (then (
        return (i32.const 1)
      )
    ))

    ;; greater than 1 means score == 5 
    local.get $radius
    f32.const 1.0
    f32.gt

    (if 
      (then (
        return (i32.const 5)
      )
    ))
    
    ;; finally, return score == 10 for being close to the centre
    (return (i32.const 10))
  )
)

