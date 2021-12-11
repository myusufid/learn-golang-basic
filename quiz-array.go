package main
import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	var lastMatch string

	fmt.Scanln(&lastMatch)
	allResults := append(results, lastMatch)
	fmt.Println(allResults)
	fmt.Println(results)


	total := 0

	for _,v := range results {

		point := 0

		if v == "w"{
			point = 3
		}else if v == "d"{
			point = 1
		}else{
			point = 0
		}

		total += point
	}

	fmt.Println(total)
}