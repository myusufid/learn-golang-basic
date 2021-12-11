package main

import "fmt"

func main()  {
	type NoKTP string
	type Married bool

	var noKtpYusuf NoKTP = "82738178273812"
	fmt.Println(noKtpYusuf)

	var MarriedStatus Married = true
	fmt.Println(MarriedStatus)

}
