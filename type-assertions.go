package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	//var result interface{} = random()
	//fmt.Println(result)
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	var result interface{} = random()
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("unknown")
	}
}
