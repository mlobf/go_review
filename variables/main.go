package main

import (
	//"bufio"
	//"container/list"
	"fmt"
	//"slices"
	//"slices"
	//"os"
)

// Ver como funciona objetos iteraveis.
// Ate agora eu tenho string.
// Vou ver listas, tuplas, dicionarios.
// Listas do Python é = a Arrays e Slices no Go.
// Arrays no Go tem tamanho fixo.
// Slices tem tamanho dinamico.

func criando_slices() []int {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4, 5, 6)
	return s

}

func iterando_no_slice() {
	iter := criando_slices()
	for i := 0; i < len(iter); i++ {
		fmt.Println(iter[i])
	}
}

func main() {
	//firstNumber := 2
	//secondNumber := 2
	//thirdNumber := 2
	//forthNumber := 2

	// Uso do input em Go
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("/")
	//fmt.Println("Hi there")
	//fmt.Println((reader.ReadString('\n')))
	//fmt.Println((reader.ReadString('\n')))
	//criando_slices()
	//iterando_no_slice()
}
