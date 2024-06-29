package main

import (
	"fmt"
	"unicode/utf8"
)

// simples funcao
func soma(z int, y int) int {
	return z + y
}

// controle de fluxo
func controle_fluxo(texto string) {
	if utf8.RuneCountInString(texto) >= 10 {
		fmt.Print("E maior que 10")
	} else {
		fmt.Print("E menor que 10")
	}

}

func main() {
	x := "oi"
	nome_do_cliente := "Marcos Leme"
	fmt.Println("Ola mundo  ", x)
	fmt.Println(soma(10, 10))
	fmt.Println("------")
	controle_fluxo(nome_do_cliente)
}
