package ttype

import (
	"gopower/convert"
)

//	ValueBool is a internal bool type object
type ValueBool struct {
	ValueBasic
	val bool
}

//	NewValueBool return a new ValueBool object with val interface
func NewValueBool(val interface{}) (valueBool *ValueBool, err error) {
	valueBool = &ValueBool{ValueBasic: ValueBasic{valueType: TypeBool, priority: PriorityBool}}

	valueBool.val, err = convert.ToBool(val)
	if err != nil {
		return nil, err
	}

	return valueBool, nil
}

//	NewValueBoolIf return interface type
func NewValueBoolIf(val interface{}) (valIf ValueIf, err error) {
	return NewValueBool(val)
}

//	Get ValueBool Value in interface mode
func (p *ValueBool) Value() (interface{}, error) {
	return p.val, nil
}

//	ValueBool Equal Method
func (p *ValueBool) Equal(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToBool()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val == rhsVal.val)
}

//	ValueBool to Int64
func (p *ValueBool) ToInt64() (*ValueInt64, error) {
	return NewValueInt64(p.val)
}

//	ValueBool to Bool
func (p *ValueBool) ToBool() (*ValueBool, error) {
	return NewValueBool(p.val)
}

//	ValueBool to Float64
func (p *ValueBool) ToFloat64() (*ValueFloat64, error) {
	return NewValueFloat64(p.val)
}
