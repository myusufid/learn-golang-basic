package main

import "fmt"

func main() {
	runApplication(10)
}

func logging() {
	fmt.Println("Selesai memanggil func logging")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result ", result)
}
