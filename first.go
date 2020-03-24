package main

import (
	"fmt"
	"math/rand"
	"time"

func main() {
	fmt.Println("Let us play a number guessing game!")
	
	//Generate a random number
	source := rand.NewSource(time.New().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10) //generates numbers between 0 and n (10)

	var guess int
	
	for {
		fmt.Println("Kindly input your number value: ")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("Your input is too big!")
		}
		else if guess < secretNumber {
			fmt.Println("Your input is too small!")
		}
		else {
			fmt.Println("You have won!")
			break
		}
	}
	
}