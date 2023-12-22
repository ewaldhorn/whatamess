"
" This function takes any remark and returns Bob's response.
"
function! Response(remark) abort

  " your solution goes here

  let s:remark = trim(a:remark)
  let strlength = len(s:remark)
  let lastChar = s:remark[strlength-1]
  let isQuestion = lastChar == '?'
  let isEmpty = len(trim(s:remark)) == 0
  let hasLetters = empty(matchstr(s:remark, '[a-zA-Z]')) == 0
  let isYelling = s:remark == toupper(s:remark) && hasLetters

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

