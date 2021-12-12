package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for i:=0; i < data.Len(); i++{
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	fmt.Println(*data)

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
