package ttype

import "github.com/woudX/gopower/convert"

//	ValueInt64 is a internal int64 type object
type ValueInt64 struct {
	ValueBasic
	val int64
}

//	NewValueInt64 return a new ValueInt64 object with val interface
func NewValueInt64(val interface{}) (valueInt64 *ValueInt64, err error) {
	valueInt64 = &ValueInt64{ValueBasic: ValueBasic{valueType: TypeInt64, priority: PriorityInt64}}

	valueInt64.val, err = convert.ToInt64(val)
	if err != nil {
		return nil, err
	}

	return valueInt64, nil
}

//	NewValueInt64If return interface type
func NewValueInt64If(val interface{}) (valIf ValueIf, err error) {
	return NewValueInt64(val)
}

//	Get ValueInt64 Value in interface mode
func (p *ValueInt64) Value() (interface{}, error) {
	return p.val, nil
}

//	ValueInt64 Equal Method
func (p *ValueInt64) Equal(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToInt64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val == rhsVal.val)
}

//	ValueInt64 Larger Method
func (p *ValueInt64) Larger(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToInt64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val > rhsVal.val)
}

//	ValueInt64 LargerEqual Method
func (p *ValueInt64) LargerEqual(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToInt64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val >= rhsVal.val)
}

//	ValueInt64 Less Method
func (p *ValueInt64) Less(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToInt64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val < rhsVal.val)
}

//	ValueInt64 LessEqual Method
func (p *ValueInt64) LessEqual(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToInt64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val <= rhsVal.val)
}

//	ValueInt64 Add Method
func (p *ValueInt64) Add(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToInt64()
	if err != nil {
		return nil, err
	}

	return NewValueInt64(p.val + rhsVal.val)
}

//	ValueInt64 to Int64
func (p *ValueInt64) ToInt64() (*ValueInt64, error) {
	return NewValueInt64(p.val)
}

//	ValueInt64 to Bool
func (p *ValueInt64) ToBool() (*ValueBool, error) {
	return NewValueBool(p.val)
}

//	ValueInt64 to Float64
func (p *ValueInt64) ToFloat64() (*ValueFloat64, error) {
	return NewValueFloat64(p.val)
}

//	ValueInt64 to Interface
func (p*ValueInt64) ToInterface() (interface{}, error) {
	return p.val, nil
}
