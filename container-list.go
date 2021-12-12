package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("M")
	data.PushBack("Yusuf")
	data.PushBack("M Yusuf")

	fmt.Println(data.Front().Prev())
	fmt.Println(data.Back().Next())

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value)
	}
}
