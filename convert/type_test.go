package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	caseList := []struct {
		Input  interface{}
		Expect int
	}{
		{
			Input:  -3,
			Expect: -3,
		},
		{
			Input:  3.14,
			Expect: 3,
		},
		{
			Input:  3.9,
			Expect: 3,
		},
		{
			Input:  true,
			Expect: 1,
		},
		{
			Input:  "313",
			Expect: 313,
		},
	}

	for _, caseItem := range caseList {
		realVal, err := ToInt(caseItem.Input)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, realVal)
	}

}

func TestToInt64(t *testing.T) {
	caseList := []struct {
		Input  interface{}
		Expect int64
	}{
		{
			Input:  -3,
			Expect: -3,
		},
		{
			Input:  3.14,
			Expect: 3,
		},
		{
			Input:  3.9,
			Expect: 3,
		},
		{
			Input:  true,
			Expect: 1,
		},
		{
			Input:  "313",
			Expect: 313,
		},
	}

	for _, caseItem := range caseList {
		realVal, err := ToInt64(caseItem.Input)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, realVal)
	}
}

func TestToBool(t *testing.T) {
	caseList := []struct {
		Input  interface{}
		Expect bool
	}{
		{
			Input:  -3,
			Expect: true,
		},
		{
			Input:  0,
			Expect: false,
		},
		{
			Input:  3.9,
			Expect: true,
		},
		{
			Input:  true,
			Expect: true,
		},
		{
			Input:  "false",
			Expect: false,
		},
	}

	for _, caseItem := range caseList {
		realVal, err := ToBool(caseItem.Input)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, realVal)
	}
}

func TestToFloat64(t *testing.T) {
	caseList := []struct {
		Input  interface{}
		Expect float64
	}{
		{
			Input:  -3,
			Expect: -3,
		},
		{
			Input:  0,
			Expect: 0,
		},
		{
			Input:  3.9,
			Expect: 3.9,
		},
		{
			Input:  true,
			Expect: 1,
		},
		{
			Input:  "3.14",
			Expect: 3.14,
		},
	}

	for _, caseItem := range caseList {
		realVal, err := ToFloat64(caseItem.Input)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, realVal)
	}
}

func TestToString(t *testing.T) {
	caseList := []struct {
		Input  interface{}
		Expect string
	}{
		{
			Input:  -3,
			Expect: "-3",
		},
		{
			Input:  0,
			Expect: "0",
		},
		{
			Input:  3.9,
			Expect: "3.9",
		},
		{
			Input:  true,
			Expect: "true",
		},
		{
			Input:  "3.14",
			Expect: "3.14",
		},
	}

	for _, caseItem := range caseList {
		realVal, err := ToString(caseItem.Input)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, realVal)
	}
}

func TestToInterfaceSlice(t *testing.T) {
	caseList := []struct {
		Input  interface{}
		Expect []interface{}
	}{
		{
			Input:  []int{1, 2, 3, 4, 5},
			Expect: []interface{}{1, 2, 3, 4, 5},
		},
		{
			Input:  []float64{1.32, 2.2, 3, 44.22, -5123.22},
			Expect: []interface{}{1.32, 2.2, float64(3), 44.22, -5123.22},
		},
		{
			Input: []interface{}{"adbs", 123, -23.123, 23323.122556, false},
			Expect: []interface{}{"adbs", 123, -23.123, 23323.122556, false},
		},
	}

	for _, caseItem := range caseList {
		realVal, err := ToInterfaceSlice(caseItem.Input)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, realVal)
	}
}
