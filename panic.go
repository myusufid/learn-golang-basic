package main

import "fmt"

func main() {
	runApp(false)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error : ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Applikasi ERROR")
	}
	fmt.Println("Aplikasi berjalana")
}
