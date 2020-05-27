package main

import(
	"math/rand"
	"fmt"
	"time"
	"bufio"
	"os"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var numberToGuess int
	numberToGuess = rand.Intn(100)
	var inputString string
	fmt.Println("input number I guessed")
        inputString , _ = bufio.NewReader(os.Stdin).ReadString('\n')
        fmt.Println(numberToGuess)  
        fmt.Println(inputString)
}
