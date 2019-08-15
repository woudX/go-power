package parammap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParamMap(t *testing.T) {
	paramMap := NewParamMap()

	paramMap.Set("int_val", 12345).
		Set("float_val", -233.233).
		Set("string_val", "test_string").
		Set("bool_val", false).
		Set("interface_val", 123)

	valInt, err := paramMap.GetInt("int_val")
	assert.Nil(t, err)
	assert.Equal(t, 12345, valInt)

	valInt64, err := paramMap.GetInt64("int_val")
	assert.Nil(t, err)
	assert.Equal(t, int64(12345), valInt64)

	valFloat64, err := paramMap.GetFloat64("float_val")
	assert.Nil(t, err)
	assert.Equal(t, -233.233, valFloat64)

	valString, err := paramMap.GetString("string_val")
	assert.Nil(t, err)
	assert.Equal(t, "test_string", valString)

	valBool, err := paramMap.GetBool("bool_val")
	assert.Nil(t, err)
	assert.Equal(t, false, valBool)

	valInterface , err := paramMap.GetInterface("interface_val")
	assert.Nil(t, err)
	assert.Equal(t, 123, valInterface)

	//	TryGet Method
	assert.Nil(t, paramMap.TryGet("int_val", &valInt))
	assert.Equal(t, 12345, valInt)
}
