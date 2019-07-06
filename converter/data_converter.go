package converter

import (
	"fmt"
	"gopower/constant"
	"gopower/powerr"
	"math"
	"reflect"
	"strconv"
)

const (
	FromType = "from_type"
	ToType   = "to_type"
)

//	Convert interface to int64
func ToInt64(srcData interface{}) (int64, error) {
	switch reflect.TypeOf(srcData).Kind() {
	case reflect.Int:
		return int64(srcData.(int)), nil
	case reflect.Int8:
		return int64(srcData.(int8)), nil
	case reflect.Int16:
		return int64(srcData.(int16)), nil
	case reflect.Int32:
		return int64(srcData.(int32)), nil
	case reflect.Int64:
		return int64(srcData.(int64)), nil
	case reflect.Uint:
		return int64(srcData.(uint)), nil
	case reflect.Uint8:
		return int64(srcData.(uint8)), nil
	case reflect.Uint16:
		return int64(srcData.(uint16)), nil
	case reflect.Uint32:
		return int64(srcData.(uint32)), nil
	case reflect.Uint64:
		return int64(srcData.(uint64)), nil
	case reflect.String:
		return strconv.ParseInt(srcData.(string), 10, 64)
	case reflect.Float32:
		return int64(srcData.(float32)), nil
	case reflect.Float64:
		return int64(srcData.(float64)), nil
	case reflect.Bool:
		if srcData.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, powerr.New(powerr.ErrNotSupportConvert).StoreKV(FromType, reflect.TypeOf(srcData).Kind()).
			StoreKV(ToType, reflect.Int64)
	}
}

//	Convert interface to bool
func ToBool(srcData interface{}) (bool, error) {
	switch reflect.TypeOf(srcData).Kind() {
	case reflect.Int:
		return (srcData.(int)) != 0, nil
	case reflect.Int8:
		return (srcData.(int8)) != 0, nil
	case reflect.Int16:
		return (srcData.(int16)) != 0, nil
	case reflect.Int32:
		return (srcData.(int32)) != 0, nil
	case reflect.Int64:
		return int64(srcData.(int64)) != 0, nil
	case reflect.Uint:
		return int64(srcData.(uint)) != 0, nil
	case reflect.Uint8:
		return int64(srcData.(uint8)) != 0, nil
	case reflect.Uint16:
		return int64(srcData.(uint16)) != 0, nil
	case reflect.Uint32:
		return int64(srcData.(uint32)) != 0, nil
	case reflect.Uint64:
		return int64(srcData.(uint64)) != 0, nil
	case reflect.String:
		return strconv.ParseBool(srcData.(string))
	case reflect.Float32:
		return math.Abs(float64(srcData.(float32))) >= constant.EPSILON, nil
	case reflect.Float64:
		return math.Abs(srcData.(float64)) >= constant.EPSILON, nil
	case reflect.Bool:
		return srcData.(bool), nil
	default:
		return false, powerr.New(powerr.ErrNotSupportConvert).StoreKV(FromType, reflect.TypeOf(srcData).Kind()).
			StoreKV(ToType, reflect.Bool)
	}
}


//	Convert interface to float64
func ToFloat64(srcData interface{}) (float64, error) {
	switch reflect.TypeOf(srcData).Kind() {
	case reflect.Int:
		return float64(srcData.(int)), nil
	case reflect.Int8:
		return float64(srcData.(int8)), nil
	case reflect.Int16:
		return float64(srcData.(int16)), nil
	case reflect.Int32:
		return float64(srcData.(int32)), nil
	case reflect.Int64:
		return float64(srcData.(int64)), nil
	case reflect.Uint:
		return float64(srcData.(uint)), nil
	case reflect.Uint8:
		return float64(srcData.(uint8)), nil
	case reflect.Uint16:
		return float64(srcData.(uint16)), nil
	case reflect.Uint32:
		return float64(srcData.(uint32)), nil
	case reflect.Uint64:
		return float64(srcData.(uint64)), nil
	case reflect.String:
		return strconv.ParseFloat(srcData.(string), 64)
	case reflect.Float32:
		return float64(srcData.(float32)), nil
	case reflect.Float64:
		return srcData.(float64), nil
	case reflect.Bool:
		if srcData.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, powerr.New(powerr.ErrNotSupportConvert).StoreKV(FromType, reflect.TypeOf(srcData).Kind()).
			StoreKV(ToType, reflect.Float64)
	}
}

//	Convert interface to string
func ToString(srcData interface{}) (string, error) {
	switch reflect.TypeOf(srcData).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
			return fmt.Sprintf("%v", srcData), nil
	default:
		return "",powerr.New(powerr.ErrNotSupportConvert).StoreKV(FromType, reflect.TypeOf(srcData).Kind()).
			StoreKV(ToType, reflect.String)
	}
}
