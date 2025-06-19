package main

import "core:fmt"

main :: proc() {
    for h in 0..=23 {
        fmt.printf("Hour: %d ->", h)

        if h >= 4 && h < 12 {
            fmt.println("Good morning")
        } else if h >= 12 && h < 17 {
            fmt.println("Good afternoon")
        } else if h >= 17 && h < 23 {
            fmt.println("Good evening")
        } else {
            fmt.println("Good night")
        }
    }
}
