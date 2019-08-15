package reflector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunctionName(t *testing.T) {
	assert.Contains(t, GetFunctionName(TestGetFunctionName), "reflector.TestGetFunctionName")
}

func TestSetVal(t *testing.T) {
	//	Normal Check
	var outValInt int
	assert.Nil(t, SetVal(3, &outValInt))
	assert.Equal(t, 3, outValInt)

	//	Ptr Check
	var valIntPtr = new(int)
	*valIntPtr = -132
	var outValIntPtr *int
	assert.Nil(t, SetVal(valIntPtr, &outValIntPtr))
	assert.Equal(t, valIntPtr, outValIntPtr)

	//	Struct Check
	type structTest struct {
		IntVal int
	}
	var outValStruct structTest

	assert.Nil(t, SetVal(structTest{3}, &outValStruct))
	assert.Equal(t, structTest{3}, outValStruct)
}
