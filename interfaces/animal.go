package main

import "fmt"

// Definindo a interface Animal
type Animal interface {
	Speak() string
}

// Tipo Dog que implementa a interface Animal
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

// Tipo Cat que implementa a interface Animal
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Tipo Cow que implementa a interface Animal
type Cow struct {
	Name string
}

func (cw Cow) Speak() string {
	return "Moo!"
}

// Função que recebe um Animal e chama seu método Speak
func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	// Criando instâncias de diferentes animais
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	cow := Cow{Name: "Bessie"}

	// Chamando a função MakeAnimalSpeak com diferentes tipos que implementam a interface Animal
	MakeAnimalSpeak(dog) // Output: Woof!
	MakeAnimalSpeak(cat) // Output: Meow!
	MakeAnimalSpeak(cow) // Output: Moo!
}
