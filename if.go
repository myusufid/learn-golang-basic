package main

import "fmt"

func main()  {
	var name = "Yusuf"

	if name == "Yusuf" {
		fmt.Println("Hello Yusuf")
	}else if name == "Dedi"{
		fmt.Println("Hello Dedi")
	}else{
		fmt.Println("Hello [unknown]")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	}else{
		fmt.Println("Nama sudah benar")
	}
}
