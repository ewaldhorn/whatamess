"
" This function takes any remark and returns Bob's response.
"
function! Response(remark) abort

  " your solution goes here

  let strlength = len(a:remark)
  let lastChar = a:remark[strlength-1]
  let isQuestion = lastChar == '?'
  
  if isQuestion
    return 'Sure.'
  endif

  return 'Whatever.'
endfunction

