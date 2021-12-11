package main

import "fmt"

func main()  {
	const firstName string = "M"
	const lastName = "Yusuf"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		year = 2020
		month = 12
	)

	fmt.Println(year)
	fmt.Println(month)
}
