package ttype

import (
	"github.com/woudX/gopower/constant"
	"github.com/woudX/gopower/convert"
	"math"
)

//	ValueFloat64 is a internal float64 type object
type ValueFloat64 struct {
	ValueBasic
	val float64
}

//	NewValueFloat64 return a new ValueInt64 object with val interface
func NewValueFloat64(val interface{}) (valueFloat64 *ValueFloat64, err error) {
	valueFloat64 = &ValueFloat64{ValueBasic: ValueBasic{valueType: TypeFloat64, priority: PriorityFloat64}}

	valueFloat64.val, err = convert.ToFloat64(val)
	if err != nil {
		return nil, err
	}

	return valueFloat64, nil
}

//	NewValueFloat64If return interface type
func NewValueFloat64If(val interface{}) (valIf ValueIf, err error) {
	return NewValueFloat64(val)
}

//	Get ValueFloat64 Value in interface mode
func (p *ValueFloat64) Value() (interface{}, error) {
	return p.val, nil
}

//	ValueFloat64 Equal Method
func (p *ValueFloat64) Equal(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToFloat64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(math.Abs(p.val-rhsVal.val) < constant.EPSILON)
}

//	ValueFloat64 Larger Method
func (p *ValueFloat64) Larger(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToFloat64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val - rhsVal.val > constant.EPSILON)
}

//	ValueFloat64 Larger Method
func (p *ValueFloat64) LargerEqual(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToFloat64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val - rhsVal.val > -constant.EPSILON)
}

//	ValueFloat64 Less Method
func (p *ValueFloat64) Less(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToFloat64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val - rhsVal.val < -constant.EPSILON)
}

//	ValueFloat64 LessEqual Method
func (p *ValueFloat64) LessEqual(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToFloat64()
	if err != nil {
		return nil, err
	}

	return NewValueBool(p.val - rhsVal.val < constant.EPSILON)
}

//	ValueFloat64 Add Method
func (p *ValueFloat64) Add(rhs ValueIf) (ValueIf, error) {
	rhsVal, err := rhs.ToFloat64()
	if err != nil {
		return nil, err
	}

	return NewValueFloat64(p.val + rhsVal.val)
}

//	ValueFloat64 to Int64
func (p *ValueFloat64) ToInt64() (*ValueInt64, error) {
	return NewValueInt64(p.val)
}

//	ValueFloat64 to Bool
func (p *ValueFloat64) ToBool() (*ValueBool, error) {
	return NewValueBool(p.val)
}

//	ValueFloat64 to Float64
func (p *ValueFloat64) ToFloat64() (*ValueFloat64, error) {
	return NewValueFloat64(p.val)
}

//	ValueFloat64 to Interface
func (p*ValueFloat64) ToInterface() (interface{}, error) {
	return p.val, nil
}
