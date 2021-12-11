package main

import "fmt"

func main()  {
	person := map[string]string{
		"name" : "Yusuf",
		"address" : "Magelang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))


	delete(person, "title")
	fmt.Println(person)



	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Yusuf"
	fmt.Println(book)


}
