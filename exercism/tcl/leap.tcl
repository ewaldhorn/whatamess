#!/usr/bin/env tclsh

proc isLeapYear {year} {
     if {[expr $year % 400] == 0} {
        return true
     }

    if {[expr $year % 4 ] == 0 && [expr $year % 100] != 0} {
        return true
    }

    return false
}

