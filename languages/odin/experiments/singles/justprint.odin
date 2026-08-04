package main

import "core:fmt"

// Define ANSI Escape Codes as constants
// \x1b is the 'Escape' character (ASCII 27)
RESET      :: "\x1b[0m"
RED        :: "\x1b[31m"
GREEN      :: "\x1b[32m"
YELLOW     :: "\x1b[33m"
BLUE       :: "\x1b[34m"
MAGENTA    :: "\x1b[35m"
CYAN       :: "\x1b[36m"
BOLD       :: "\x1b[1m"

main :: proc() {
    // 1. Simple color printing
    fmt.println(RED + "This text is bright red!" + RESET)
    fmt.println(GREEN + "This text is vibrant green!" + RESET)
    
    // 2. Combining styles (Bold + Yellow)
    fmt.print(BOLD + YELLOW)
    fmt.println("This is a bold yellow warning message." + RESET)
    
    // 3. Mixing colors in a single line using fmt.printf
    fmt.printf("Status: %s[ ONLINE ]%s | System: %s[ STABLE ]%s\n", GREEN, RESET, CYAN, RESET)
    
    // 4. Showing what happens if you forget to reset
    fmt.print(MAGENTA)
    fmt.println("If I don't use the RESET constant here...")
    fmt.println("everything staying after this line will stay magenta until the terminal resets!")
    fmt.print(RESET) 
}

