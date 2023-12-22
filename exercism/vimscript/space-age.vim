"
" Returns an age in local years, given a planet name and number of seconds on Earth
" Raises an exception if the name is not a valid planet in the solar system
"
function! Age(planet, seconds) abort
  " we need every planet in this list
  let s:earthSeconds = 31557600

  if a:planet == "Mercury"
    let planetSeconds = 0.2408467 * s:earthSeconds
    return a:seconds / planetSeconds
  elseif a:planet == "Venus"
    let planetSeconds = 0.61519726 * s:earthSeconds
    return a:seconds / planetSeconds
  elseif a:planet == "Earth"
    let planetSeconds = 1.0 * s:earthSeconds
    return a:seconds / planetSeconds
  elseif a:planet == "Mars"
    let planetSeconds = 1.8808158 * s:earthSeconds
    return a:seconds / planetSeconds
  elseif a:planet == "Jupiter"
    let planetSeconds = 11.862615 * s:earthSeconds
    return a:seconds / planetSeconds
  elseif a:planet == "Saturn"
    let planetSeconds = 29.447498 * s:earthSeconds
    return a:seconds / planetSeconds
  elseif a:planet == "Uranus"
    let planetSeconds = 84.016846 * s:earthSeconds
    return a:seconds / planetSeconds
  elseif a:planet == "Neptune"
    let planetSeconds = 164.79132 * s:earthSeconds
    return a:seconds / planetSeconds
  else
    throw 'not a planet'
endfunction

