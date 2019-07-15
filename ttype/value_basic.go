package ttype

import (
	"github.com/woudX/gopower/powerr"
	"github.com/woudX/gopower/reflector"
)

const (
	FuncName = "func_name"
)

const (
	PriorityLowest = iota
	PriorityBool
	PriorityString
	PriorityInt64
	PriorityFloat64
	PriorityHighest
)

//	ValueBasic is base of all other type, it contains common used data and functions
type ValueBasic struct {
	valueType ValueType
	priority  int
}

//	Basic Functions
//	Return Priority
func (p *ValueBasic) Priority() int {
	return p.priority
}

//	Return Type
func (p *ValueBasic) Type() ValueType {
	return p.valueType
}

//	Return Value
func (p *ValueBasic) Value() (interface{}, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.Value))
}

//	Operations
//	Equal
func (p *ValueBasic) Equal(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.Equal))
}

//	Larger
func (p *ValueBasic) Larger(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.Larger))
}

//	LargerEqual
func (p *ValueBasic) LargerEqual(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.LargerEqual))
}

//	Less
func (p *ValueBasic) Less(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.Less))
}

//	LessEqual
func (p *ValueBasic) LessEqual(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.LessEqual))
}

//	Converter
//	ToInt64
func (p *ValueBasic) ToInt64() (*ValueInt64, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.ToInt64))
}

//	ToFloat64
func (p *ValueBasic) ToFloat64() (*ValueFloat64, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.ToFloat64))
}

//	ToBool
func (p *ValueBasic) ToBool() (*ValueBool, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.ToBool))
}

//	ToString
func (p *ValueBasic) ToString() (*ValueString, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.ToString))
}

//	ToInterface
func (p *ValueBasic) ToInterface() (interface{}, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.ToInterface))
}
