package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "pressione o enter"

func main() {
	firstNumber := 2
	secondNumber := 5
	subtract := 7
	//var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the number")
	fmt.Println("-----------------------")
	fmt.Println(" ")
	fmt.Println("Pense em um numero entre 1 e 10", prompt)
	reader.ReadString('\n')
	fmt.Println("Multiplique o seu numero por.", firstNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Agora faça pelo numero o seu numero.", secondNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Agora divida pelo resultado pelo numero que vc havia pensado enteriormente.")
	reader.ReadString('\n')
	fmt.Println("Now subtract", subtract, prompt)
	reader.ReadString('\n')
}
