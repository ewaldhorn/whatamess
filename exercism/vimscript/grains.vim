"
" Returns the number of grains on a chessboard square given the grains on each square
" double from the previous square. 
" Throws an error if the square is below 1 or above 64.
"
" Examples:
"
"   :echo Square(16)
"   32768
"
"   :echo Square(-1)
"   E605: Exception not caught: square must be between 1 and 64
"
function! Square(number) abort
  if a:number < 1 || a:number > 64
    throw 'square must be between 1 and 64'
  endif

  " remember to subtract 1 from the total for the calculation
  return float2nr(pow(2, a:number - 1))
endfunction

"
" Returns the total number of grains for a filled chessboard
"
function! Total() abort
  return float2nr(pow(2, 64) - 1)
endfunction

