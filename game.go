package main

import(
	"math/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var numberToGuess int
	numberToGuess = rand.Intn(100)
	fmt.Println(numberToGuess)
}
