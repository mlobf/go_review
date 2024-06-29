package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//firstNumber := 2
	//secondNumber := 2
	//thirdNumber := 2
	//forthNumber := 2

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("/")
	fmt.Println("Hi there")
	reader.ReadString('\n')

}
