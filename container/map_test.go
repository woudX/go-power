package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapTryGet(t *testing.T) {
	caseList := []struct {
		Input  map[string]interface{}
		Key    string
		Expect interface{}
	}{
		{
			Input:  map[string]interface{}{"A": 1, "B": 3.114115, "C": "string"},
			Key:    "A",
			Expect: 1,
		},
		{
			Input:  map[string]interface{}{"A": 1, "B": 3.114115, "C": "string"},
			Key:    "B",
			Expect: 3.114115,
		},
		{
			Input:  map[string]interface{}{"A": 1, "B": 3.114115, "C": "string"},
			Key:    "C",
			Expect: "string",
		},
	}

	for _, caseItem := range caseList {
		var result interface{}

		assert.Nil(t, MapTryGet(caseItem.Input, caseItem.Key, &result))
		assert.Equal(t, caseItem.Expect, result)
	}
}
