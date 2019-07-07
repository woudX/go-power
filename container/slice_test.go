package container

import (
	"gopower/src/github.com/stretchr/testify/assert"
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
