package main

import(
	"math/rand"
	"fmt"
	"time"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var numberToGuess int
	numberToGuess = rand.Intn(101)
	var inputString string
	var inputNumber int
	var tries int
	for {
		tries = tries + 1
		fmt.Println("input number I guessed")
		inputString, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		inputNumber, _ = strconv.Atoi(strings.TrimSpace(inputString))
		if (inputNumber > numberToGuess) {
			fmt.Println("Number I guessed is smaller. Try again")
		} else if (inputNumber < numberToGuess) {
			fmt.Println("Number I guessed is larger. Try again")
		} else {
			fmt.Println("You made", tries, "tries")
			fmt.Println("Well done 🎉🎉🎉")
			break
		}		
	}
}
