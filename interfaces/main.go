package main

import "fmt"

// interfaces
type Describer interface {
	Describe() string
}
type Person struct {
	Name string
	Age  int
}

// Implementando o metodo do Objeto Person
func (p Person) Describe() string {
	return fmt.Sprintf("Name: %s, Age:%d", p.Name, p.Age)
}

func main() {
	fmt.Println("BEG")
	person := Person{Name: "Marcos", Age: 42}
	var d Describer = person
	fmt.Println(d.Describe())
}
