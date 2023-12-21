;; Exercism Leap Year
;; on every year that is evenly divisible by 4
;;   except every year that is evenly divisible by 100
;;     unless the year is also evenly divisible by 400
(module
     (func (export "isLeap")(param $year i32)(result i32)
        (i32.and
            ;; check if we are divisible by 4
            (i32.eq (i32.rem_u (local.get $year)(i32.const 4))(i32.const 0))
            ;; now check if we are not divisible by 100, but still by 400
            (i32.or 
                (i32.ne (i32.rem_u (local.get $year)(i32.const 100))(i32.const 0))
                (i32.eq (i32.rem_u (local.get $year)(i32.const 400))(i32.const 0))
            )
        )
        ;; function returns the last result implicitly
     )
)
