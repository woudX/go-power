package ttype

import "gopower/converter"

//	ValueBool is a internal bool type object
type ValueBool struct {
	ValueBasic
	val bool
}

//	NewValueBool return a new ValueBool object with val interface
func NewValueBool(val interface{}) (valueBool *ValueBool, err error) {
	valueBool = &ValueBool{ValueBasic: ValueBasic{valueType: TypeBool, priority:PriorityBool}}

	valueBool.val, err = converter.ToBool(val)
	if err != nil {
		return nil, err
	}

	return valueBool, nil
}

//	NewValueBoolIf return interface type
func NewValueBoolIf(val interface{}) (valIf ValueIf, err error) {
	return NewValueBool(val)
}

//	ValueBool Equal Method
func (p *ValueBool) Equal(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToBool()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val == rhsVal.val)
}

//	ValueBool to Bool
func (p *ValueBool) ToBool() (*ValueBool, error) {
	return NewValueBool(p.val)
}