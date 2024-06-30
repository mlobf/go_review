package main

import (
	"fmt"
	//"strings"
	"reflect"
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

func iteracao(variavel_iteravel string) {
	for i := 0; i < utf8.RuneCountInString(variavel_iteravel); i++ {
		fmt.Println(i)
	}
}

func printLetters(s string) {
	for _, char := range s {
		fmt.Println(string(char))
	}
}

// Função que itera e imprime cada caractere de uma string usando índice
func printLetter(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}

// Valeu o Drill para relembrar.
func iterar_letra(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}

// Essa abordagem nao funcionou.
// Me parece que o int convertido para
// string nao apareceu.
func printar_int_string(numero int) {
	fmt.Println("VAmos que vamos")
	fmt.Println(string(numero))
}

func verificar_tipagem(i string) {
	fmt.Println(reflect.TypeOf(i))
}

func print(s string) {
	fmt.Println(s)
}

func main() {
	x := "oi"
	nome_do_cliente := "Marcos Leme"
	fmt.Println("Ola mundo  ", x)
	fmt.Println(soma(10, 10))
	fmt.Println("------")
	controle_fluxo(nome_do_cliente)
	fmt.Println("")
	//iteracao(nome_do_cliente)
	//printLetters(nome_do_cliente)
	//printLetter(nome_do_cliente)
	//fmt.Printf(string(nome_do_cliente[2]))
	//iterar_letra(nome_do_cliente)
	//printar_int_string(10)
	verificar_tipagem("testes")
	print("Marcos")
}
