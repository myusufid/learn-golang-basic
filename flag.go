package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your user")
	var password *string = flag.String("password", "password", "Put your password")

	flag.Parsed()

	fmt.Println("Host", *host)
	fmt.Println("User", *user)
	fmt.Println("Password", *password)
}
