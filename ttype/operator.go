package ttype

import "github.com/woudX/gopower/powerr"

//============================= Operator Interface ===============================//

//	Operator Interface only contain a function to do operate
type OperatorIf interface {
	Operate(valueIf ...ValueIf) (ValueIf, error)
}

//============================= Operator Above ===============================//

var OpEqual = &OperatorEqual{}

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
