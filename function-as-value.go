package main

import "fmt"

func main()  {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Yusuf")
	fmt.Println(result)
}

func getGoodBye(name string) string{
	return "Good Bye " + name
}
