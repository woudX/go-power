package ttype

import (
	"gopower/powerr"
	"gopower/reflector"
)

const (
	FuncName = "func_name"
)

const (
	PriorityLowest = iota
	PriorityBool
	PriorityInt64
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

//	Operations
//	Equal
func (p *ValueBasic) Equal(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.Equal))
}

//	Larger
func (p *ValueBasic) Larger(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.Larger))
}

//	Less
func (p *ValueBasic) Less(rhs ValueIf) (ValueIf, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.Less))
}

//	Converter
//	ToInt64
func (p *ValueBasic) ToInt64() (*ValueInt64, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.ToInt64))
}

//	ToBool
func (p *ValueBasic) ToBool() (*ValueBool, error) {
	return nil, powerr.New(powerr.ErrInfoUnImplementFunction).StoreKV(FuncName, reflector.GetFunctionName(p.ToBool))
}
