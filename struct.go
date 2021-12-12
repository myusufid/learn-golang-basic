package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//func sayHi(customer Customer, name string) {
//	fmt.Println("Hello", name, "My name is", customer.Name)
//}

// struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {

	var yusuf Customer
	yusuf.Name = "Yusuf"
	yusuf.Age = 19
	yusuf.Address = "Yogyakarta"

	fmt.Println(yusuf)

	joko := Customer{
		Name:    "Joko",
		Address: "Yogyakarta",
		Age:     32,
	}
	fmt.Println(joko)
	//sayHi(joko, "Joko")

	// struct method
	joko.sayHi("Yusuf")
}
