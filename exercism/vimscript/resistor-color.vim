"
" Returns the numeric value of the passed-in resistor band color
"
" Example:
"
"   :echo ColorCode('black')
"   0
"
"   :echo ColorCode('green')
"   5
"
"   :echo ColorCode('white')
"   9
"
function! ColorCode(color) abort
  " get the position of the colour out of the Colors array
  let position = index(Colors(), a:color)
  if position < 0
    throw 'Invalid color'
  else
    return position
  endif
endfunction

"
" Returns a list of resistor color bands by color name
"
" Example:
"
"   :echo Colors()
"   ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']
"
function! Colors() abort
  " return the list of colours as an array, which is zero-indexed, and black is 0 anyway
  return ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']
endfunction

