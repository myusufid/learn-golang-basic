package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("M Yusuf", "M"))
	fmt.Println(strings.Split("M Yusuf", " "))
	fmt.Println(strings.ToLower("M Yusuf"))
	fmt.Println(strings.ToUpper("M Yusuf"))
	fmt.Println(strings.ToTitle("M Yusuf"))
	fmt.Println(strings.Trim("  M Yusuf ", " "))
	fmt.Println(strings.ReplaceAll("M M M M", "M", "S"))
}
