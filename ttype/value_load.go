package ttype

import (
	"github.com/woudX/gopower/powerr"
	"reflect"
)

//	LoadValueIfFromInterface can load data from interface, check type and create library internal ValueType
//	These types can be used to compare and calculate
func LoadValueIfFromInterface(val interface{}) (valIf ValueIf, err error) {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valIf, err = NewValueInt64(val)
	case reflect.Bool:
		valIf, err = NewValueBool(val)
	case reflect.String:
		valIf, err = NewValueString(val)
	case reflect.Float32, reflect.Float64:
		valIf, err = NewValueFloat64(val)
	default:
		err = powerr.New(powerr.ErrNotSupportValueType).StoreKV("Type", reflect.TypeOf(val).Kind()).StoreStack()
	}

	return valIf, err
}
