package main

import "fmt"

func GuessMyNumber(game Game) string {
	low := 1
	high := 360
	var guess int
	var res string

	for low <= high {
		guess = (low + high) / 2
		res = game.CheckNumber(guess)

		if res == "Correct" {
			return "Your Number was " + itoa(guess)
		} else if res == "My Number is Greater" {
			low = guess + 1
		} else if res == "My Number is Lower" {
			high = guess - 1
		} else {
			return "Your Number was " + itoa(guess)
		}
	}

	return "Number not found in range"
}

func itoa(num int) string {
	return fmt.Sprintf("%d", num)
}
