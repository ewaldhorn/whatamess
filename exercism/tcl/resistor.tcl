namespace eval resistorColor {
    
    proc colorCode {args} {
        set i [lsearch -exact [colors] $args]
        if {$i == -1} {
            error "Invalid color: $args"
        } else {
            return $i
        }
    }

    proc colors {args} {
        return [list black brown red orange yellow green blue violet grey white ]
    }
}

