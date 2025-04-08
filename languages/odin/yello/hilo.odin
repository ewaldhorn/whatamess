package main

import "core:fmt"
import "core:math/rand"

main :: proc() {
	// Set the range for the game
	min := 1
	max := 100

	// Generate a random number within the range
	secret_number := rand.int() % (max - min + 1) + min

	// Initialize the number of guesses
	num_guesses := 0

	fmt.println("Welcome to the high/low guessing game!")
	fmt.printf("I'm thinking of a number between %d and %d.\n", min, max)

	for {
		// Ask the player for their guess
		fmt.print("Enter your guess: ")
		guess := 0
		fmt.scanln(&guess)

		// Check if the guess is valid
		if guess < min || guess > max {
			fmt.println("Invalid guess! Please try again.")
			continue
		}

		// Increment the number of guesses
		num_guesses += 1

		// Check if the guess is correct
		if guess == secret_number {
			fmt.printf(" Congratulations! You found the number in %d guesses.\n", num_guesses)
			break
		} else if guess < secret_number {
			fmt.println("Higher!")
		} else {
			fmt.println("Lower!")
		}
	}
}
