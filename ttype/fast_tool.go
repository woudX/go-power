package ttype

import (
	"gopower/powerr"
	"reflect"
)

func TryGetBoolFromValueIf(valIf ValueIf) (result bool, err error) {
	valBool, ok := valIf.(*ValueBool)
	if !ok {
		return false, powerr.New(powerr.ErrNotSupportConvert).StoreKV("from", valIf.Type()).StoreKV("to", reflect.Bool)
	}

	return valBool.val, nil
}