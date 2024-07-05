package main

import (
	"fmt"
	"math/rand"
)

func while_loop() {
	i := 1000
	for i > 100 {
		i = rand.Intn(1000) + 1
		fmt.Println("i is ", i)
		if i > 100 {
			fmt.Println("i is ", i, " so loops keeps going")
		}
	}
	fmt.Println("Got ", i, " and broke out of loop")
}

func main() {
	fmt.Println("Oi mundo !!!")
	// execut a loop while i > 100
	while_loop()
}
