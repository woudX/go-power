package ttype

import "github.com/woudX/gopower/convert"

//	ValueString is a internal string type object
type ValueString struct {
	ValueBasic
	val string
}

//	NewValueString return a new NewValueString object with val interface
func NewValueString(val interface{}) (valueString *ValueString, err error) {
	valueString = &ValueString{ValueBasic: ValueBasic{valueType: TypeString, priority: PriorityString}}

	valueString.val, err = convert.ToString(val)
	if err != nil {
		return nil, err
	}

	return valueString, nil
}

//	NewValueStringIf return interface type
func NewValueStringIf(val interface{}) (valIf ValueIf, err error) {
	return NewValueString(val)
}

//	Get ValueString Value in interface mode
func (p *ValueString) Value() (interface{}, error) {
	return p.val, nil
}

//	ValueString Equal Method
func (p *ValueString) Equal(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToString()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val == rhsVal.val)
}

//	ValueString To String
func (p *ValueString) ToString() (*ValueString, error) {
	return NewValueString(p.val)
}
