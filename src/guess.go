package main

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

	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make a guess :")

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(target)

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}

	if success != true {
		fmt.Println("Target value is", target)
	}

}