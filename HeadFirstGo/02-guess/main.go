package main

// guess challenges players to guess a random number.
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. you guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. You guess was HIGH")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			fmt.Println("You guessed it in: ", guesses, "guesses Well Done.")
			break
		}
	}

	if !success {
		fmt.Println("Sorry you didn't guess my number. It was: ", target)
	}
	fmt.Println("Thank you for playing would you like another go? ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	if input == "Yes" {
		main()
	} else if input == "No" {
		fmt.Println("Thank you for playing please play again soon")
	} else if input == "yes" {
		main()
	} else if input == "no" {
		fmt.Println("Thank you for playing please play again soon")
	}
}
