package main

import "fmt"

func main()  {
	a,b,c := getFullNameName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func getFullNameName() (firstName string, middleName string, lastName string){
	firstName = "M"
	middleName = " "
	lastName = "Yusuf"

	return
}
