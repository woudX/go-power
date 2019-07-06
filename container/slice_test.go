package container

import (
	"gopower/src/github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceFind(t *testing.T) {
	testCase := []struct {
		Input []interface{}
		Target interface{}
		Expect int
	}{
		{
			Input: []interface{}{true, 2, 3456, 233},
			Target: 3456,
			Expect: 2,
		},
		{
			Input: []interface{}{true, 1, 3456, 233, false},
			Target: true,
			Expect: 1,
		},
		{
			Input: []interface{}{true, 1, 3456, 233},
			Target: -33,
			Expect: -1,
		},
	}

	for _, caseItem := range testCase {
		pos, err := SliceFind(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}

func TestSliceFindString(t *testing.T) {
	testCase := []struct {
		Input []string
		Target string
		Expect int
	}{
		{
			Input: []string{"aaa", "b1c2d3", "dddee122"},
			Target: "dddee122",
			Expect: 2,
		},
		{
			Input: []string{"aaa", "b1c2d3", "dddee122"},
			Target: "aaa",
			Expect: 0,
		},
		{
			Input: []string{"aaa", "b1c2d3", "dddee122", ""},
			Target: "ccc",
			Expect: -1,
		},
		{
			Input: []string{"aaa", "b1c2d3", "dddee122", ""},
			Target: "",
			Expect: 3,
		},
	}

	for _, caseItem := range testCase {
		pos, err := SliceFindString(caseItem.Input, caseItem.Target)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}