package main

import "fmt"

func main()  {
	sum(1,2,3,4,5)
	sum(2,6)

	s := []int{32,51,2,2,3}
	sum(s...)
}

func sum(nums ...int){
	total := 0
	for _, value := range nums {
		total += value
	}
	fmt.Println(total)
}

