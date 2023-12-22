"
" This function takes any remark and returns Bob's response.
"
function! Response(remark) abort

  " your solution goes here

  let strlength = len(a:remark)
  let lastChar = a:remark[strlength-1]
  let isQuestion = lastChar == '?'
  let isEmpty = len(trim(a:remark)) == 0
  let isYelling = a:remark == toupper(a:remark)

  if isEmpty
    return 'Fine. Be that way!'
  elseif isYelling && isQuestion
    return 'Calm down, I know what I''m doing!'
  elseif isYelling
    return 'Whoa, chill out!'
  elseif isQuestion
    return 'Sure.'
  else
    return 'Whatever.'
  endif
endfunction

