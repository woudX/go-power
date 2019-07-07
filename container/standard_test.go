package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	testCase := []struct {
		Input []interface{}
		Target interface{}
		Expect int
	}{
		{
			Input: []interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 22,
			Expect: 3,
		},
		{
			Input: []interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: "string-A",
			Expect: 4,
		},
		{
			Input: []interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: -33,
			Expect: -1,
		},
		{
			Input: []interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.1,
			Expect: 7,
		},
		{
			Input: []interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111},
			Target: 3.111,
			Expect: 9,
		},
	}

	for _, caseItem := range testCase {
		pos, err := Find(caseItem.Target, caseItem.Input...)
		assert.Nil(t, err)
		assert.Equal(t, caseItem.Expect, pos)
	}
}