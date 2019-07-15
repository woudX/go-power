package ttype

import "github.com/woudX/gopower/powerr"

//============================= Operator Interface ===============================//

//	Operator Interface only contain a function to do operate
type OperatorIf interface {
	Operate(valueIf ...ValueIf) (ValueIf, error)
}

//============================= Operator Above ===============================//

var (
	OpEqual       = &OperatorEqual{}
	OpLarger      = &OperatorLarger{}
	OpLargerEqual = &OperatorLargerEqual{}
	OpLess        = &OperatorLess{}
	OpLessEqual   = &OperatorLessEqual{}

	OpAdd = &OperatorAdd{}
)

//	Operator Equal
type OperatorEqual struct{}

//	Operator Equal Operator
func (p *OperatorEqual) Operate(valueIf ...ValueIf) (result ValueIf, err error) {

	// need two valueIf
	if len(valueIf) != 2 {
		return nil, powerr.New(powerr.ErrNotEnoughParams).StoreKV("param_len", len(valueIf))
	}

	leftVal := valueIf[0]
	rightVal := valueIf[1]

	//	check priority and do operator
	if leftVal.Priority() >= rightVal.Priority() {
		return leftVal.Equal(rightVal)
	} else {
		return rightVal.Equal(leftVal)
	}
}

//	Operator Larger
type OperatorLarger struct{}

//	Operator Larger Operator
func (p *OperatorLarger) Operate(valueIf ...ValueIf) (result ValueIf, err error) {

	// need two valueIf
	if len(valueIf) != 2 {
		return nil, powerr.New(powerr.ErrNotEnoughParams).StoreKV("param_len", len(valueIf))
	}

	leftVal := valueIf[0]
	rightVal := valueIf[1]

	//	check priority and do operator
	if leftVal.Priority() >= rightVal.Priority() {
		return leftVal.Larger(rightVal)
	} else {
		return rightVal.Less(leftVal)
	}
}

//	Operator Less
type OperatorLess struct{}

//	Operator Less Operator
func (p *OperatorLess) Operate(valueIf ...ValueIf) (result ValueIf, err error) {

	// need two valueIf
	if len(valueIf) != 2 {
		return nil, powerr.New(powerr.ErrNotEnoughParams).StoreKV("param_len", len(valueIf))
	}

	leftVal := valueIf[0]
	rightVal := valueIf[1]

	//	check priority and do operator
	if leftVal.Priority() >= rightVal.Priority() {
		return leftVal.Less(rightVal)
	} else {
		return rightVal.Larger(leftVal)
	}
}

//	Operator LargerEqual
type OperatorLargerEqual struct{}

//	Operator LargerEqual Operator
func (p *OperatorLargerEqual) Operate(valueIf ...ValueIf) (result ValueIf, err error) {

	// need two valueIf
	if len(valueIf) != 2 {
		return nil, powerr.New(powerr.ErrNotEnoughParams).StoreKV("param_len", len(valueIf))
	}

	leftVal := valueIf[0]
	rightVal := valueIf[1]

	//	check priority and do operator
	if leftVal.Priority() >= rightVal.Priority() {
		return leftVal.LargerEqual(rightVal)
	} else {
		return rightVal.Less(leftVal)
	}
}

//	Operator LessEqual
type OperatorLessEqual struct{}

//	Operator LessEqual Operator
func (p *OperatorLessEqual) Operate(valueIf ...ValueIf) (result ValueIf, err error) {

	// need two valueIf
	if len(valueIf) != 2 {
		return nil, powerr.New(powerr.ErrNotEnoughParams).StoreKV("param_len", len(valueIf))
	}

	leftVal := valueIf[0]
	rightVal := valueIf[1]

	//	check priority and do operator
	if leftVal.Priority() >= rightVal.Priority() {
		return leftVal.LessEqual(rightVal)
	} else {
		return rightVal.Larger(leftVal)
	}
}

//	Operator Add
type OperatorAdd struct{}

//	Operator Add Operator
func (p *OperatorAdd) Operate(valueIf ...ValueIf) (result ValueIf, err error) {

	// need two valueIf
	if len(valueIf) != 2 {
		return nil, powerr.New(powerr.ErrNotEnoughParams).StoreKV("param_len", len(valueIf))
	}

	leftVal := valueIf[0]
	rightVal := valueIf[1]

	//	check priority and do operator
	if leftVal.Priority() >= rightVal.Priority() {
		return leftVal.Add(rightVal)
	} else {
		return rightVal.Add(leftVal)
	}
}
