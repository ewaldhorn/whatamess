set earthYearInSeconds 31557600.0
set planetOrbits {
    Neptune 164.79132 
}
proc onEarth {seconds} {
    return [expr $seconds / $::earthYearInSeconds]
}
proc onMercury {seconds} {
    set planetSeconds [expr {0.2408467 * $::earthYearInSeconds}]
    return [expr {1.0 * $seconds / $planetSeconds}]
}
proc onVenus {seconds} {    
    set planetSeconds [expr {0.61519726 * $::earthYearInSeconds}]
    return [expr {1.0 * $seconds / $planetSeconds}]
}
proc onMars {seconds} {    
    set planetSeconds [expr {1.8808158 * $::earthYearInSeconds}]
    return [expr {1.0 * $seconds / $planetSeconds}]
}
proc onJupiter {seconds} {    
    set planetSeconds [expr {11.862615 * $::earthYearInSeconds}]
    return [expr {1.0 * $seconds / $planetSeconds}]
}
proc onSaturn {seconds} {    
    set planetSeconds [expr {29.447498 * $::earthYearInSeconds}]
    return [expr {1.0 * $seconds / $planetSeconds}]
}
proc onUranus {seconds} {    
    set planetSeconds [expr {84.016846 * $::earthYearInSeconds}]
    return [expr {1.0 * $seconds / $planetSeconds}]
}
proc onNeptune {seconds} {    
    set planetSeconds [expr {164.79132 * $::earthYearInSeconds}]
    return [expr {1.0 * $seconds / $planetSeconds}]
}
proc onSun {seconds} {error "not a planet"}
# extra credit: generate the procedures programmatically
