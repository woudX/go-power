package ttype

import (
	"fmt"
	"gopower/gp_error"
	"gopower/gp_reflect"
)

type BasicType struct {
	ValueData interface{} `json:"value_data"`
	ValueType ValueType
}

func Type(p *BasicType) ValueType {
	return p.ValueType
}

//	Operations
//	Equal
func (p *BasicType) Equal(rhs ValueIf) (ValueIf, error) {
	return nil, gp_error.New(fmt.Sprint(gp_error.ErrInfoUnImplementFunction, gp_reflect.GetFunctionName(p.Equal)))
}

//	Larger
func (p *BasicType) Larger(rhs ValueIf) (ValueIf, error) {
	return nil, gp_error.New(fmt.Sprintf(gp_error.ErrInfoUnImplementFunction, gp_reflect.GetFunctionName(p.Larger)))
}

//	Less
func (p *BasicType) Less(rhs ValueIf) (ValueIf, error) {
	return nil, gp_error.New(fmt.Sprintf(gp_error.ErrInfoUnImplementFunction, gp_reflect.GetFunctionName(p.Less)))
}
