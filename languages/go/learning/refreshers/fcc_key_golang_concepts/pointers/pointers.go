package main

import "fmt"

func main() {
    var num int = 10

    var ptr *int = &num

    fmt.Println("Value of num:", num)      
    fmt.Println("Pointer address:", ptr)   
    fmt.Println("Value at pointer:", *ptr)

    *ptr = 20
    fmt.Println("Updated value of num:", num) 
}
