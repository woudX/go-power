package ttype

import (
	"gopower/powerr"
	"reflect"
)

//	LoadValueFromInterface can load data from interface, check type and create library internal ValueType
//	These types can be used to compare and calculate
func LoadValueFromInterface(val interface{}) (valIf ValueIf, err error) {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valIf, err = NewValueInt64(val)
	case reflect.Bool:
		valIf, err = NewValueBool(val)
	default:
		err = powerr.New(powerr.ErrNotSupportValueType).StoreKV("Type", reflect.TypeOf(val).Kind()).StoreStack()
	}

	return valIf, err
}
