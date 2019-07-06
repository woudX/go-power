package main

import (
	"fmt"
	"gopower/container"
)

func main() {
	input := []interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111}

	pos, err := container.SliceFind(input, "string-A")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(pos)
}