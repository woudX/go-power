package main

import (
	"fmt"
	"gopower/container"
	"gopower/convert"
)

var demoInput = []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111}
var demoMultiInput = []interface{}{-3, 544, true, 22, "string-A", 123, 22, 22, "str-B", 3.1, -23.4, 22, 3.111}
var demoStringInput = []string{"aaa", "string-B", "string-A"}

func main() {
	pos, err := container.FindInSlice(convert.MustToInterfaceSlice(demoStringInput), "string-A")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err := container.RemoveFirstSlice(demoInput, "string-A")
	fmt.Printf("result=%v, err=%v\n", result, err)

	result, err = container.RemoveFirstSlice(result, 3.1)
	fmt.Printf("result=%v, err=%v\n", result, err)

	result, err = container.RemoveFirstSlice(result, -23.4)
	fmt.Printf("result=%v, err=%v\n", result, err)

	result, err = container.RemoveFirstSlice(result, 3.1)
	fmt.Printf("result=%v, err=%v\n", result, err)

	result, err = container.RemoveFromSlice(demoMultiInput, 22)
	fmt.Printf("result=%v, err=%v\n", result, err)

	fmt.Printf("%v", pos)
}
