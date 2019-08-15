package container

import (
	"fmt"
	"github.com/woudX/gopower/powerr"
	"github.com/woudX/gopower/reflector"
	"reflect"
)

// GoPower provide a series of c-like method to operate map


//	MapTryGet accept any kind of map and key, return the value in map with out type, for example:
//
//	var outInt int
//	inMap := map[string]int{"int_val": 3}
//
//	container.MapTryGet(inMap, "int_val", &outInt)
//
//   method support all kinds of data type including struct
func MapTryGet(inMap interface{}, key interface{}, out interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = powerr.New("MapTryGet() cause panic").StoreKV("recover_info", r)
		}
	}()

	inMapValue := reflect.ValueOf(inMap)
	keyValue := reflect.ValueOf(key)

	if inMapValue.IsNil() || inMapValue.Kind() != reflect.Map  {
		return powerr.New(fmt.Sprintf("innvalid inMap type : %v", inMapValue.Kind()))
	}

	//	Get value from map
	targetVal := inMapValue.MapIndex(keyValue)
	if !targetVal.IsValid() {
		out = nil
		return powerr.New("key not exist").StoreKV("key", keyValue.Interface())
	}

	return reflector.SetVal(targetVal.Interface(), out)
}


//	MapTryGetWithString method use reflect to get value with key and return with out param
func MapTryGetWithString(inMap map[string]interface{}, key string, out interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = powerr.New("MapTryGetWithString() cause panic").StoreKV("recover_info", r)
		}
	}()

	realVal, exist := inMap[key]
	if !exist {
		return powerr.New("key not exist").StoreKV("key", key)
	}

	return reflector.SetVal(realVal, out)
}
