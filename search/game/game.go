package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	fmt.Println("I am thinking of a number between 0-99 (inclusive).")
	fmt.Println("Type in your guess. I will tell you 'higher' or 'lower' until you are correct!")
	answer := rand.Intn(100)
	br := bufio.NewReader(os.Stdin)
	tries := 0
	for {
		tries++
		fmt.Printf("\nAttempt #%d\n", tries)
		br.Reset(os.Stdin)
		fmt.Print("> ")
		buf, err := br.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		}
		input := buf[:len(buf)-1]
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Println("Input must be an integer")
			tries--
			continue
		}
		if guess == answer {
			fmt.Printf("Correct! I was thinking of %d!\n", answer)
			fmt.Printf("Guesses: %d\n", tries)
			return
		}
		if guess < answer {
			fmt.Println("Higher!")
		} else {
			fmt.Println("Lower!")
		}
	}
}
