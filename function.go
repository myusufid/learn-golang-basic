package main

import "fmt"

func main()  {
	hello := getHello("Yusuf")
	fmt.Println(hello)

	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

// function return value
func getHello(name string) string{
	return "Hello " + name
}

// return multiple
//func getFullName() (string, string){
//	return "M", "Yusuf"
//}