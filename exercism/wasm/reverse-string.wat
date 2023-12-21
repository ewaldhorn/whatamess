(module
  (memory (export "mem") 1)
 
  ;;
  ;; Reverse a string
  ;;
  ;; @param {i32} offset - The offset of the input string in linear memory
  ;; @param {i32} length - The length of the input string in linear memory
  ;;
  ;; @returns {(i32,i32)} - The offset and length of the reversed string in linear memory
  ;;
  (func (export "reverseString") (param $offset i32) (param $length i32) (result i32 i32)
    ;; keep track of start, end and swap character
    (local $start i32)
    (local $end i32)
    (local $swap i32)

    (if (i32.gt_u (local.get $length)(i32.const 1))
    (then 

    ;; set the start of the string
    (local.set $start (local.get $offset))

    ;; set end of the string - remember it is 0 indexed, so subtract 1
    (local.set $end (i32.add (local.get $offset)(local.get $length)))
    (local.set $end (i32.sub (local.get $end)(i32.const 1)))

    (loop $lopp
      ;; swap start and end characters
      ;; first we save the start value in swap
      (i32.store8 (local.get $swap)(i32.load8_u (local.get $start)))
      ;; now we save the end value in start
      (i32.store8 (local.get $start)(i32.load8_u (local.get $end)))
      ;; now we put the swap value into end
      (i32.store8 (local.get $end)(i32.load8_u (local.get $swap)))

      ;; increment start
      (local.set $start (i32.add (local.get $start)(i32.const 1)))

      ;; decrement end
      (local.set $end (i32.sub (local.get $end)(i32.const 1)))

      ;; check if we need to loop again - as long as start is less than end
      (br_if $lopp (i32.lt_s (local.get $start)(local.get $end)))
    )
    ))

    (return (local.get $offset) (local.get $length))
  )
)

