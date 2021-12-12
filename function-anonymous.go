package main

import "fmt"

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("yusuf", blackList)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
