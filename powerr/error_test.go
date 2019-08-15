package powerr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowError_Error(t *testing.T) {
	caseList := []struct {
		Input  string
		Expect string
	}{
		{
			Input:  "this is error info",
			Expect: "this is error info",
		},
	}

	for _, caseItem := range caseList {
		err := New(caseItem.Input)
		assert.NotNil(t, err)

		assert.Equalf(t, caseItem.Expect, err.Error(), fmt.Sprintf("%v", caseItem))
	}
}

func TestPowError_StoreKV(t *testing.T) {
	caseList := []struct {
		Input  string
		KV     map[string]interface{}
		Expect string
	}{
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_int": 1,
			},
			Expect: "this is error info||key_int=1",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_float": -12.333,
			},
			Expect: "this is error info||key_float=-12.333",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_bool": false,
			},
			Expect: "this is error info||key_bool=false",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_string": "test_string",
			},
			Expect: "this is error info||key_string=test_string",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_struct": struct {
					A int
					B float64
				}{
					333,
					-231.23,
				},
			},
			Expect: "this is error info||key_struct={333 -231.23}",
		},
	}

	for _, caseItem := range caseList {
		err := New(caseItem.Input)

		for key, val := range caseItem.KV {
			err.StoreKV(key, val)
		}

		assert.NotNil(t, err)
		assert.Equalf(t, caseItem.Expect, err.Error(), fmt.Sprintf("%v", caseItem))
	}
}


func TestPowError_ErrorWithSep(t *testing.T) {
	caseList := []struct {
		Input  string
		KV     map[string]interface{}
		Expect string
	}{
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_int": 1,
			},
			Expect: "this is error info&&key_int=1",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_float": -12.333,
			},
			Expect: "this is error info&&key_float=-12.333",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_bool": false,
			},
			Expect: "this is error info&&key_bool=false",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_string": "test_string",
			},
			Expect: "this is error info&&key_string=test_string",
		},
		{
			Input: "this is error info",
			KV: map[string]interface{}{
				"key_struct": struct {
					A int
					B float64
				}{
					333,
					-231.23,
				},
			},
			Expect: "this is error info&&key_struct={333 -231.23}",
		},
	}

	for _, caseItem := range caseList {
		err := New(caseItem.Input)

		for key, val := range caseItem.KV {
			err.StoreKV(key, val)
		}

		assert.NotNil(t, err)
		assert.Equalf(t, caseItem.Expect, err.ErrorWithSep("&&"), fmt.Sprintf("%v", caseItem))
	}
}

func TestPowError_StoreStack(t *testing.T) {
	err := New("stack_test").StoreStack()
	assert.NotNil(t, err)

	assert.Contains(t, err.Error(), "/powerr/error_test.go")
}