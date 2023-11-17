package main

import "fmt"

//A funcao obrigatoria main nao pode receber parametros....
func main() {
	//var whatToSay string
	//whatToSay = "Hello Word Go"
	whatToSay := "Hello World Go"
	sayHelloWord(whatToSay)
}

func sayHelloWord(whatToSay string) {
	fmt.Println(whatToSay)
}
