package mathex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/woudX/gopower/convert"
)

func TestSum(t *testing.T) {
	caseList := []struct{
		Input []interface{}
		Expect interface{}
	} {
		{
			Input: []interface{}{20, 50, -30, 111},
			Expect: 151,
		},
		{
			Input: []interface{}{33.11, 558.123, -234.912, 22.1123, 44.1234, -500, 314},
			Expect: 236.5567,
		},
	}

	for _, caseItem := range caseList {
		result, err := Sum(caseItem.Input...)
		assert.Nil(t, err)

		cmpResult, err := Compare(convert.MustToFloat64(result), convert.MustToFloat64(caseItem.Expect))
		assert.Nil(t, err)
		assert.Equal(t, 0, cmpResult)
	}
}
