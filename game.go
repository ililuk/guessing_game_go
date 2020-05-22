package main

import(
	"math/rand"
         "fmt"
)

func main() {
	var numberToGuess int
	numberToGuess = rand.Intn(100)
	fmt.Println(numberToGuess)
}
