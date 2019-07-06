package mathex

import "gopower/ttype"

//	Compare two interface value and return result, return -1 if lVal larger than rVal, return 0 if equal and
//	return 1 if lVal smaller than rVal
func Compare(lVal interface{}, rVal interface{}) (cmpResult int, err error) {
	lValIf, err := ttype.LoadValueIfFromInterface(lVal)
	if err != nil {
		return 0, err
	}

	rValIf, err := ttype.LoadValueIfFromInterface(rVal)
	if err != nil {
		return 0, err
	}

	//	check if value equal
	equalResult, err := ttype.OpEqual.Operate(lValIf, rValIf)
	if err != nil {
		return 0, err
	}

	equalResultVal, err := ttype.TryGetBoolFromValueIf(equalResult)
	if err != nil {
		return 0, err
	}

	if equalResultVal {
		return 0, nil
	}

	//	if not equal, then check which is larger one

}
