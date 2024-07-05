package main

import "fmt"

func while_true() {
	i := true
	for i == true {
		fmt.Print("oi dentro do while true")
	}
}

func main() {
	fmt.Println("Oi ")
	while_true()
}
