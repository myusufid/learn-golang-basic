package main

import "fmt"

type HasName interface {
	GetName() string
}

// function ini memiliki kontrak dengan HasName
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// function milik struct Person
func (person Person) GetName() string {
	return person.Name
}

func main() {
	var yusuf Person
	yusuf.Name = "M Yusuf"

	SayHello(yusuf)
}
