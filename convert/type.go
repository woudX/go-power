package convert

import (
	"fmt"
	"github.com/woudX/gopower/constant"
	"github.com/woudX/gopower/powerr"
	"math"
	"reflect"
	"strconv"
)

const (
	FromType = "from_type"
	ToType   = "to_type"
)

//	Convert interface to int
func ToInt(srcData interface{}) (int, error) {
	switch reflect.TypeOf(srcData).Kind() {
	case reflect.Int:
		return srcData.(int), nil
	case reflect.Int8:
		return int(srcData.(int8)), nil
	case reflect.Int16:
		return int(srcData.(int16)), nil
	case reflect.Int32:
		return int(srcData.(int32)), nil
	case reflect.Int64:
		return int(srcData.(int64)), nil
	case reflect.Uint:
		return int(srcData.(uint)), nil
	case reflect.Uint8:
		return int(srcData.(uint8)), nil
	case reflect.Uint16:
		return int(srcData.(uint16)), nil
	case reflect.Uint32:
		return int(srcData.(uint32)), nil
	case reflect.Uint64:
		return int(srcData.(uint64)), nil
	case reflect.String:
		result, err := strconv.ParseInt(srcData.(string), 10, 64)
		return int(result), err
	case reflect.Float32:
		return int(srcData.(float32)), nil
	case reflect.Float64:
		return int(srcData.(float64)), nil
	case reflect.Bool:
		if srcData.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, powerr.New(powerr.ErrNotSupportConvert).StoreKV(FromType, reflect.TypeOf(srcData).Kind()).
			StoreKV(ToType, reflect.Int)
	}
}

//	Must ToInt version
func MustToInt(srcData interface{}) int {
	result, err := ToInt(srcData)
	if err != nil {
		panic("converter: MustToInt(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
}

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

//	Must ToInt64 version
func MustToInt64(srcData interface{}) int64 {
	result, err := ToInt64(srcData)
	if err != nil {
		panic("converter: MustToInt64(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
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

//	Must ToBool version
func MustToBool(srcData interface{}) bool {
	result, err := ToBool(srcData)
	if err != nil {
		panic("converter: MustToBool(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
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

//	Must ToFloat64 version
func MustToFloat64(srcData interface{}) float64 {
	result, err := ToFloat64(srcData)
	if err != nil {
		panic("converter: MustToFloat64(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
}

//	Convert interface to string
func ToString(srcData interface{}) (string, error) {
	switch reflect.TypeOf(srcData).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return fmt.Sprintf("%v", srcData), nil
	default:
		return "", powerr.New(powerr.ErrNotSupportConvert).StoreKV(FromType, reflect.TypeOf(srcData).Kind()).
			StoreKV(ToType, reflect.String)
	}
}

//	Must ToString version
func MustToString(srcData interface{}) string {
	result, _ := ToString(srcData)
	return result
}

//	Convert interface to interface list, this function only works on slice
func ToInterfaceSlice(srcData interface{}) ([]interface{}, error) {
	if reflect.TypeOf(srcData).Kind() != reflect.Slice {
		return nil, powerr.New(powerr.ErrNotSupportConvert).StoreKV(FromType, reflect.TypeOf(srcData).Kind()).
			StoreKV(ToType, reflect.Slice)
	}

	//	Create src and dst slice
	srcSlice := reflect.ValueOf(srcData)
	dstSlice := reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, srcSlice.Len())

	//	Copy src to dst
	for idx := 0; idx < srcSlice.Len(); idx++ {
		dstSlice = reflect.Append(dstSlice, srcSlice.Index(idx))
	}

	return dstSlice.Interface().([]interface{}), nil
}

//	MustToInterfaceSlice version is like ToInterfaceSlice but panic if convert failed
func MustToInterfaceSlice(srcData interface{}) []interface{} {
	result, err := ToInterfaceSlice(srcData)
	if err != nil {
		panic("converter: MustToInterfaceSlice(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
}
