package main

import "fmt"

func main() {
	counter := 0
	name := "Yusuf"

	increment := func() {
		fmt.Println("Increment")
		counter++
		name := "iamyusuf"
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
