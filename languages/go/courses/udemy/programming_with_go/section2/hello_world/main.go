package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Alas, here we go again!")
	problem_2()
	problem_3()
	problem_4()
	problem_5()
	problem_6()
	problem_7()
	problem_8()
}

func problem_2() {
	fmt.Println("Errr, it's Ewald?")
}

func problem_3() {
	fmt.Println("I'm Ewald.\nI do things with computers.\nSometimes it works.")
}

func problem_4() {
	for i := range 10 {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
}

func problem_5() {
	fmt.Printf("2 ** 11 is %.0f\n", math.Pow(2.0, 11.0))
}

func problem_6() {
	for i := range 1000 {
		fmt.Println(i + 1)
	}
}

func problem_7() {
	fmt.Println("Your lucky number is:", rand.Intn(11))
}

func problem_8() {
	fmt.Println("The current date is", time.Now())
}
