(module
  ;;
  ;; Return the number of steps needed to reach 1 in the Collatz conjecture.
  ;;
  ;; @param {i32} number - The number to start from.
  ;;
  ;; @returns {i32} - The number of steps needed to reach 1.
  ;;
  (func (export "steps") (param $number i32) (result i32)
    (local $tmp i32)
    (local $count i32)
    ;; first check for a positive number
    ;; if we don't have one, abort the process early
    (if (i32.lt_s (local.get $number)(i32.const 1))(then (return (i32.const -1))))

    ;; set a sane value for $tmp, our counter
    (local.set $count (i32.const 0))

    ;; if the number is greater than 1, start the process
    (if
      (i32.gt_s (local.get $number)(i32.const 1))
      (then
        ;; we can work with this number
        ;; Take any positive integer n. If n is even, divide n by 2 to get n / 2. 
        ;; If n is odd, multiply n by 3 and add 1 to get 3n + 1. 
        ;; Repeat the process indefinitely. The conjecture states that no matter which number you start with, 
        ;; you will always reach 1 eventually.
        (local.set $tmp (local.get $number))

        ;; we need to loop until we get to 1
        (loop $lopp
          ;; increment our step counter
          (local.set $count (i32.add (local.get $count)(i32.const 1)))

          ;; now see if we have an even number
          (if (i32.eqz (i32.rem_u (local.get $tmp)(i32.const 2)))
            (then
              ;; is even, divide by 2
              (local.set $tmp (i32.div_s (local.get $tmp)(i32.const 2)))
            )
            (else
              ;; not even, so times 3 and add 1
              (local.set $tmp (i32.mul (local.get $tmp)(i32.const 3)))
              (local.set $tmp (i32.add (local.get $tmp)(i32.const 1)))
            )
          )

          ;; branch to the loop if we have not reached 1 yet
          (br_if $lopp (i32.ne (local.get $tmp)(i32.const 1)))
        )
      )
    )

    (return (local.get $count))
  )
)
