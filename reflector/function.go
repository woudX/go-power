package reflector

import (
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

