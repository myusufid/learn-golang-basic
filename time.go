package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Print(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Second())

	utc := time.Date(2020,10,10,10,10,109,10,time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}
