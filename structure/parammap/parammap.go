package parammap

import (
	"github.com/woudX/gopower/powerr"
	"github.com/woudX/gopower/reflector"
)

//	ParamMap is a simple package of map[string]interface{}, this struct is purpose to
//	make map get/set easy and safety, you don't need to do exist-checking and convert-checking,
//	just use get/set to get the data you want
type ParamMap struct {
	innerMap map[string]interface{}
}

//	NewParamMap return a ParamMap pointer to use
func NewParamMap() *ParamMap {
	return &ParamMap{
		innerMap: make(map[string]interface{}),
	}
}

//	Set (key, val) pair to ParamMap, if key exist just override old data
func (pm *ParamMap) Set(key string, val interface{}) *ParamMap {
	pm.innerMap[key] = val
	return pm
}

//	TryGet method use reflect to get value with key and return with out param, for example:
//
//	var outInt int
// 	pm.Set("int_val", 3)
//	pm.TryGet("int_val", &outInt)
//
//   method support all kinds of data type including struct,
func (pm *ParamMap) TryGet(key string, out interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = powerr.New("ParamMap TryGet() cause panic").StoreKV("recover_info", r)
		}
	}()

	realVal, exist := pm.innerMap[key]
	if !exist {
		return powerr.New("key not exist").StoreKV("key", key)
	}

	return reflector.SetVal(realVal, out)
}
