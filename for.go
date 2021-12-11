package main

import "fmt"

func main()  {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// for range
	slice := []string{"M", "Yusuf"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Yusuf", "iamyusuf"}

	for index, name := range names {
		fmt.Println("Index ", index, " = ", name)
	}

	// variable _ tidak digunakan
	for _, name := range names{
		fmt.Println(name)
	}




}
