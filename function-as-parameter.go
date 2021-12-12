package main

import "fmt"

func main() {
	sayHelloWithFilter("Yusuf", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}

//func sayHelloWithFilter(name string, filter func(string) string) {
//	nameFiltered := filter(name)
//	fmt.Println("Hello", nameFiltered)
//}

// Filter : function sayHelloWithFilter dengan type function
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
