package container

import (
	"github.com/stretchr/testify/assert"
	"github.com/woudX/gopower/convert"
	"github.com/woudX/gopower/mathex"
	"testing"
)

func TestFindInSliceInterface(t *testing.T) {
	testCase := []struct {
		Input  []interface{}
		Target interface{}
		Expect int
	}{
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 22,
			Expect: 3,
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: "string-A",
			Expect: 4,
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: -33,
			Expect: -1,
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.1,
			Expect: 7,
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.111,
			Expect: 9,
		},
	}

	for _, caseItem := range testCase {
		pos, err := FindInSlice(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestFindLastInSlice(t *testing.T) {
	testCase := []struct {
		Input  []interface{}
		Target interface{}
		Expect int
	}{
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 22, 3.111},
			Target: 22,
			Expect: 9,
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, "string-A", -23.4, 3.111},
			Target: "string-A",
			Expect: 8,
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: -33,
			Expect: -1,
		},
		{
			Input:  []interface{}{-3, 544, true, 3.10, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.1,
			Expect: 8,
		},
		{
			Input:  []interface{}{-3, 544, 3.111, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.111,
			Expect: 10,
		},
	}

	for _, caseItem := range testCase {
		pos, err := FindLastInSlice(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestFindInSliceCmp (t *testing.T) {
	testCase := []struct {
		Input       []interface{}
		CompareFunc func(lVal interface{}, rVal interface{}) (result int, err error)
		CompareVal  interface{}
		Expect      interface{}
	}{
		{
			Input: []interface{}{-3, 5, 2, 7 ,11, -4, 3, 2, 2, 11},
			CompareFunc: func(customVal interface{}, containerVal interface{}) (result int, err error) {

				//	Find larger equal than custom val
				result, _ = mathex.Compare(containerVal, customVal)

				if result >= 0 {
					return 1, nil
				} else {
					return 0, nil
				}
			},
			CompareVal: 3,
			Expect: 1,
		},
	}

	for _, caseItem := range testCase {
		result, err := FindInSliceCmp(caseItem.Input, caseItem.CompareVal, caseItem.CompareFunc)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, result)
	}
}


func TestFindLastInSliceCmp (t *testing.T) {
	testCase := []struct {
		Input       []interface{}
		CompareFunc func(lVal interface{}, rVal interface{}) (result int, err error)
		CompareVal  interface{}
		Expect      interface{}
	}{
		{
			Input: []interface{}{-3, 5, 2, 7 ,11, -4, 3, 2, 2, 11},
			CompareFunc: func(customVal interface{}, containerVal interface{}) (result int, err error) {

				//	Find larger equal than custom val
				result, _ = mathex.Compare(containerVal, customVal)

				if result >= 0 {
					return 1, nil
				} else {
					return 0, nil
				}
			},
			CompareVal: 3,
			Expect: 9,
		},
	}

	for _, caseItem := range testCase {
		result, err := FindLastInSliceCmp(caseItem.Input, caseItem.CompareVal, caseItem.CompareFunc)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, result)
	}
}

func TestFindInSliceIf(t *testing.T) {
	testCase := []struct {
		Input  []interface{}
		Func   ifFunc
		Expect int
	}{
		{
			Input: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 22, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if intVal, ok := val.(int); ok {
					return convert.ToInt(intVal%2 == 1)
				}

				return 0, nil
			},
			Expect: 5,
		},
		{
			Input: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, "string-A", -23.4, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if intVal, ok := val.(int); ok {
					return convert.ToInt(intVal%2 == 0)
				}

				return 0, nil
			},
			Expect: 1,
		},
		{
			Input: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if _, ok := val.(string); ok {
					return 1, nil
				}

				return 0, nil
			},
			Expect: 4,
		},
		{
			Input: []interface{}{-3, 544, true, 3.10, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if _, ok := val.(float64); ok {
					return 1, nil
				}

				return 0, nil
			},
			Expect: 3,
		},
	}

	for _, caseItem := range testCase {
		result, err := FindInSliceIf(caseItem.Input, caseItem.Func)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, result)
	}
}

func TestFindInSliceString(t *testing.T) {
	testCase := []struct {
		Input  []string
		Target string
		Expect int
	}{
		{
			Input:  []string{"aaa", "b1c2d3", "dddee122"},
			Target: "dddee122",
			Expect: 2,
		},
		{
			Input:  []string{"aaa", "b1c2d3", "dddee122"},
			Target: "aaa",
			Expect: 0,
		},
		{
			Input:  []string{"aaa", "b1c2d3", "dddee122", ""},
			Target: "ccc",
			Expect: -1,
		},
		{
			Input:  []string{"aaa", "b1c2d3", "dddee122", ""},
			Target: "",
			Expect: 3,
		},
	}

	for _, caseItem := range testCase {
		pos, err := FindInSliceString(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestFindInSliceInt64(t *testing.T) {
	testCase := []struct {
		Input  []int64
		Target int64
		Expect int
	}{
		{
			Input:  []int64{113, 444, -321, 5, 0},
			Target: 444,
			Expect: 1,
		},
		{
			Input:  []int64{113, 444, -321, 5, 0},
			Target: -321,
			Expect: 2,
		},
		{
			Input:  []int64{113, 444, -321, 5, 0},
			Target: -3,
			Expect: -1,
		},
	}

	for _, caseItem := range testCase {
		pos, err := FindInSliceInt64(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestFindInSliceInt(t *testing.T) {
	testCase := []struct {
		Input  []int
		Target int
		Expect int
	}{
		{
			Input:  []int{113, 444, -321, 5, 0},
			Target: 444,
			Expect: 1,
		},
		{
			Input:  []int{113, 444, -321, 5, 0},
			Target: -321,
			Expect: 2,
		},
		{
			Input:  []int{113, 444, -321, 5, 0},
			Target: -3,
			Expect: -1,
		},
	}

	for _, caseItem := range testCase {
		pos, err := FindInSliceInt(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestFindInSliceFloat32(t *testing.T) {
	testCase := []struct {
		Input  []float32
		Target float32
		Expect int
	}{
		{
			Input:  []float32{113.33, 444.44, -321.21, 5.0, 5, 0},
			Target: 444.44,
			Expect: 1,
		},
		{
			Input:  []float32{113.33, 444.44, -321.21, 5.0, 5, 5.0001, 0},
			Target: 5,
			Expect: 3,
		},
		{
			Input:  []float32{113.33, 444.44, -321.21, 5.0, 5, 5.0001, 0},
			Target: 5.0001,
			Expect: 5,
		},
		{
			Input:  []float32{113.33, 444.44, -321.21, 5.0, 5, 5.0001, 0},
			Target: -321.2,
			Expect: -1,
		},
	}

	for _, caseItem := range testCase {
		pos, err := FindInSliceFloat32(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestFindInSliceFloat64(t *testing.T) {
	testCase := []struct {
		Input  []float64
		Target float64
		Expect int
	}{
		{
			Input:  []float64{113.33, 444.44, -321.21, 5.0, 5, 0},
			Target: 444.44,
			Expect: 1,
		},
		{
			Input:  []float64{113.33, 444.44, -321.21, 5.0, 5, 5.0001, 0},
			Target: 5,
			Expect: 3,
		},
		{
			Input:  []float64{113.33, 444.44, -321.21, 5.0, 5, 5.0001, 0},
			Target: 5.0001,
			Expect: 5,
		},
		{
			Input:  []float64{113.33, 444.44, -321.21, 5.0, 5, 5.0001, 0},
			Target: -321.2,
			Expect: -1,
		},
	}

	for _, caseItem := range testCase {
		pos, err := FindInSliceFloat64(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestRemoveFromFirstSlice(t *testing.T) {
	testCase := []struct {
		Input  []interface{}
		Target interface{}
		Expect []interface{}
	}{
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 22, 3.111},
			Target: 22,
			Expect: []interface{}{-3, 544, true, "string-A", 123, "str-B", 3.1, -23.4, 22, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, "string-A", -23.4, 3.111},
			Target: "string-A",
			Expect: []interface{}{-3, 544, true, 22, 123, "str-B", 3.1, "string-A", -23.4, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: -33,
			Expect: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, true, 3.10, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.1,
			Expect: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, 3.111, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.111,
			Expect: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
		},
	}

	for _, caseItem := range testCase {
		newSlice, err := RemoveFirstFromSlice(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, newSlice)
	}
}

func TestRemoveFromSliceSlice(t *testing.T) {
	testCase := []struct {
		Input  []interface{}
		Target interface{}
		Expect []interface{}
	}{
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 22, 3.111},
			Target: 22,
			Expect: []interface{}{-3, 544, true, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, "string-A", -23.4, 3.111},
			Target: "string-A",
			Expect: []interface{}{-3, 544, true, 22, 123, "str-B", 3.1, -23.4, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: -33,
			Expect: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, true, 3.10, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.1,
			Expect: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", -23.4, 3.111},
		},
		{
			Input:  []interface{}{-3, 544, 3.111, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.111,
			Expect: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4},
		},
		{
			Input:  []interface{}{-3, -3, -3, -3, -3, -3},
			Target: -3,
			Expect: []interface{}{},
		},
	}

	for _, caseItem := range testCase {
		newSlice, err := RemoveFromSlice(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, newSlice)
	}
}

func TestRemoveFromSliceCmp(t *testing.T) {
	testCase := []struct {
		Input       []interface{}
		CompareFunc func(lVal interface{}, rVal interface{}) (result int, err error)
		CompareVal  interface{}
		Expect      []interface{}
	}{
		{
			Input: []interface{}{-3, 5, 2, 7 ,11, -4, 3, 2, 2, 11},
			CompareFunc: func(customVal interface{}, containerVal interface{}) (result int, err error) {

				//	Remove larger equal than custom val
				result, _ = mathex.Compare(containerVal, customVal)

				if result >= 0 {
					return 1, nil
				} else {
					return 0, nil
				}
			},
			CompareVal: 3,
			Expect: []interface{}{-3, 2, -4, 2, 2},
		},
	}

	for _, caseItem := range testCase {
		result, err := RemoveFromSliceCmp(caseItem.Input, caseItem.CompareVal, caseItem.CompareFunc)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, result)
	}
}

func TestRemoveFromSliceIf(t *testing.T) {
	testCase := []struct {
		Input  []interface{}
		Func   ifFunc
		Expect []interface{}
	}{
		{
			Input: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 22, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if intVal, ok := val.(int); ok {
					return convert.ToInt(intVal%2 == 1)
				}

				return 0, nil
			},
			Expect: []interface{}{-3, 544, true, 22, "string-A", "str-B", 3.1, -23.4, 22, 3.111},
		},
		{
			Input: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, "string-A", -23.4, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if intVal, ok := val.(int); ok {
					return convert.ToInt(intVal%2 == 0)
				}

				return 0, nil
			},
			Expect: []interface{}{-3, true, "string-A", 123, "str-B", 3.1, "string-A", -23.4, 3.111},
		},
		{
			Input: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if _, ok := val.(string); ok {
					return 1, nil
				}

				return 0, nil
			},
			Expect: []interface{}{-3, 544, true, 22, 123, 3.1, -23.4, 3.111},
		},
		{
			Input: []interface{}{-3, 544, true, 3.10, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Func: func(val interface{}) (result int, err error) {
				if _, ok := val.(float64); ok {
					return 1, nil
				}

				return 0, nil
			},
			Expect: []interface{}{-3, 544, true, 22, "string-A", 123, "str-B"},
		},
	}

	for _, caseItem := range testCase {
		result, err := RemoveFromSliceIf(caseItem.Input, caseItem.Func)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, result)
	}
}
