package main

import (
	"fmt"
	"time"
)

func simpleChannelExercise() {
	channel := make(chan int)

	go func() { channel <- 1 }()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}

func moreChannelExercises() {
	channel := make(chan int, 2)

	channel <- 1
	channel <- 2

	go func() { channel <- 3 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}

func channelExerciseThree() {
	one := make(chan int)
	two := make(chan int)

	for {
		select {
		case o := <-one:
			fmt.Println("one:", o)
		case t := <-two:
			fmt.Println("two:", t)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout reached, aborting.")
			return
		}
	}
}

func main() {
	fmt.Printf("\nChannels.\n")

	simpleChannelExercise()
	moreChannelExercises()
	channelExerciseThree()

	fmt.Println()
}
