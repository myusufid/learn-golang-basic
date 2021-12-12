package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil{
		fmt.Println(name)
	}else{
		fmt.Print("error", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
