package mathex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	caseList := []struct {
		Left   interface{}
		Right  interface{}
		Expect interface{}
	}{
		{
			Left:   333,
			Right:  333,
			Expect: 0,
		},
		{
			Left:   3,
			Right:  true,
			Expect: 1,
		},
		{
			Left:   333,
			Right:  "333",
			Expect: 0,
		},
		{
			Left:   333,
			Right:  333.0,
			Expect: 0,
		},
		{
			Left:   333,
			Right:  333.1,
			Expect: -1,
		},
		{
			Left:   333,
			Right:  332.999,
			Expect: 1,
		},
		{
			Left:   "string",
			Right:  "strinf",
			Expect: 1,
		},
		{
			Left:   "string",
			Right:  "strinh",
			Expect: -1,
		},
		{
			Left:   "string",
			Right:  "strin",
			Expect: 1,
		},
		{
			Left:   "334",
			Right:  335,
			Expect: -1,
		},
	}

	for _, caseItem := range caseList {
		result, err := Compare(caseItem.Left, caseItem.Right)
		assert.Nil(t, err)

		assert.Equalf(t, caseItem.Expect, result, fmt.Sprintf("%v", caseItem))
	}
}
