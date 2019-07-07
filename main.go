package main

import (
	"fmt"
	"gopower/container"
	"gopower/convert"
)

var demoInput = []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111}
var demoMultiInput = []interface{}{-3, 544, true, 22, "string-A", 123, 22, 22, "str-B", 3.1, -23.4, 22, 3.111}
var demoStringInput = []string{"aaa", "string-B", "string-A"}
var demoIntInput = []int{2, 4, 6, 9, 10, 11}

func main() {
	pos, err := container.FindInSlice(convert.MustToInterfaceSlice(demoStringInput), "string-A")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(pos)

	result, err := container.FindInSliceIf(convert.MustToInterfaceSlice(demoIntInput), func(val interface{}) (result int, err error) {
		return convert.ToInt(val.(int) % 2 == 1)
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)


	result2, err := container.RemoveFromSliceIf(convert.MustToInterfaceSlice(demoIntInput), func(val interface{}) (result int, err error) {
		return convert.ToInt(val.(int) % 2 == 1)
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result2)
}
