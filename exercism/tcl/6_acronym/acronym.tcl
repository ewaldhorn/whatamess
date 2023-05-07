proc abbreviate {phrase} {
    # remove characters _',
    set work [string map {"_" ""} $phrase]

    # string must be in uppercase
    set upper [string toupper $work]

    # word splitter on - and space
    set workList [split $upper { -}]

    set result {}

    foreach x $workList {
        lappend result [string index $x 0]
    }

    return [join $result ""]
}

