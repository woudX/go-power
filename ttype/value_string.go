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

//	ValueString Larger Method
func (p *ValueString) Larger(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToString()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val > rhsVal.val)
}

//	ValueString LargerEqual Method
func (p *ValueString) LargerEqual(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToString()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val >= rhsVal.val)
}

//	ValueString Less Method
func (p *ValueString) Less(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToString()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val < rhsVal.val)
}

//	ValueString LessEqual Method
func (p *ValueString) LessEqual(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToString()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val <= rhsVal.val)
}

//	ValueString To String
func (p *ValueString) ToString() (*ValueString, error) {
	return NewValueString(p.val)
}

//	ValueString To Int64
func (p *ValueString) ToInt64() (*ValueInt64, error) {
	return NewValueInt64(p.val)
}

//	ValueString To Float64
func (p *ValueString) ToFloat64() (*ValueFloat64, error) {
	return NewValueFloat64(p.val)
}

//	ValueString To Boolean
func (p *ValueString) ToBool() (*ValueBool, error) {
	return NewValueBool(p.val)
}

//	ValueString to Interface
func (p*ValueString) ToInterface() (interface{}, error) {
	return p.val, nil
}