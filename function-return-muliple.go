package main

import "fmt"

func main()  {
	firstname, _ := getFullName()
	fmt.Println(firstname)
}

func getFullName() (string, string){
	return "M", "Yusuf"
}
