(module
  (memory (export "mem") 1)

  ;;
  ;; Calculate the hamming distance between two strings.
  ;;
  ;; @param {i32} firstOffset - The offset of the first string in linear memory.
  ;; @param {i32} firstLength - The length of the first string in linear memory.
  ;; @param {i32} secondOffset - The offset of the second string in linear memory.
  ;; @param {i32} secondLength - The length of the second string in linear memory.
  ;;
  ;; @returns {i32} - The hamming distance between the two strings or -1 if the
  ;;                  strings are not of equal length.
  ;;
  (func (export "compute") 
    (param $firstOffset i32) (param $firstLength i32) (param $secondOffset i32) (param $secondLength i32) (result i32)
      ;; define local variables
      (local $tmp i32)    
      (local $position i32)

      ;; set sane defaults for variables
      (local.set $tmp (i32.const -1))
      (local.set $position (i32.const 0))

      ;; check for empty strands first
      (if
        (i32.and
          (i32.eqz (local.get $firstLength))
          (i32.eqz (local.get $secondLength)))
        (then
          ;; strands are empty, no difference
          (return (i32.const 0))
        )
      )

      ;; let's compare the string lengths
      (i32.ne (local.get $firstLength)(local.get $secondLength))
      (if
        (then
          ;; strands are a different length, it's no go 
          (return (i32.const -1))
        )
      )

      ;; we know the strings are not empty and are of equal length
      ;; so lets compare them char for char

      ;; first set our difference counter to zero
      (local.set $tmp (i32.const 0))

      ;; loop over the strands using the position counter
      (loop
        (if 
          ;; if the two characters are not the same in the same position, add one to the tmp distance variable
          (i32.ne 
            (i32.load8_u (i32.add (local.get $firstOffset)(local.get $position)))
            (i32.load8_u (i32.add (local.get $secondOffset)(local.get $position)))
          )
          (then(local.set $tmp (i32.add (local.get $tmp)(i32.const 1))))
        )
        ;; break out of the loop if we are no longer less than the first length
        (br_if 0 (i32.lt_u (local.tee $position (i32.add (local.get $position)(i32.const 1)))(local.get $firstLength)))
      )
    
      (return (local.get $tmp))
  )
)

