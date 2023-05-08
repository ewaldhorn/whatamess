oo::class create HighScores {
    variable scores

    constructor {} {
        variable scores
        set scores {}
    }
    
    method addScores {args} {
        variable scores
        lappend scores {*}$args
    }
    
    method scores {} {
        variable scores
        return $scores
    }
    
    method latest {} {
        lindex [my scores] end
    }
    
    method personalBest {} {
        lindex [my sorted] 0
    }
    
    method topThree {} {
        lrange [my sorted] 0 2
    }
    
    method sorted {} {
        lsort -integer -decreasing [my scores]
    }
}

