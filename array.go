package main

import "fmt"

func main()  {
	var names [2]string

	names[0] = "M"
	names[1] = "Yusuf"


	var values = [3]int{
		2,3,5,
	}
	fmt.Println(values)

	fmt.Println(len(values))
}
