package reflector

import (
	"github.com/woudX/gopower/powerr"
	"reflect"
	"runtime"
)

//	Get function names
func GetFunctionName(funcPtr interface{}) string {
	refVal := reflect.ValueOf(funcPtr)

	//	Check if Valid
	if refVal.IsValid() {
		return runtime.FuncForPC(refVal.Pointer()).Name()
	} else {
		return ""
	}
}

//	SetVal require a val and a referenceOut
func SetVal(val interface{}, referenceOut interface{}) error {
	//	outVal must Ptr
	outVal := reflect.ValueOf(referenceOut)
	if outVal.Kind() != reflect.Ptr || outVal.IsNil() {
		return powerr.New("out type must be references value and not nil").StoreKV("kind", outVal.Kind())
	}

	if !outVal.Elem().CanSet() {
		return powerr.New("out type can't set by reflect.Set()").StoreKV("kind", outVal.Kind())
	}

	outVal.Elem().Set(reflect.ValueOf(val))
	return nil
}
