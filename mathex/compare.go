package mathex

import (
	"github.com/woudX/gopower/ttype"
)

//	Compare two interface value and return result, return -1 if lVal larger than rVal, return 0 if equal and
//	return 1 if lVal smaller than rVal
func Compare(lVal interface{}, rVal interface{}) (result int, err error) {
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
	cmpResult, err := ttype.OpLarger.Operate(lValIf, rValIf)
	if err != nil {
		return 0, err
	}
	cmpResultVal, err := ttype.TryGetBoolFromValueIf(cmpResult)
	if err != nil {
		return 0, err
	}

	if cmpResultVal {
		return 1, nil
	} else {
		return -1, nil
	}
}

//	Max return the maximum of params with interface mode
func Max(params ...interface{}) (result interface{}, err error) {
	if len(params) <= 0 {
		return nil, nil
	}

	maxVal, err := ttype.LoadValueIfFromInterface(params[0])
	if err != nil {
		return nil, err
	}

	//	compare and get maxVal
	for _, item := range params {
		nowVal, err := ttype.LoadValueIfFromInterface(item)
		if err != nil {
			return nil, err
		}

		isLargerEqual, err := ttype.OpLargerEqual.Operate(maxVal, nowVal)
		if err != nil {
			return nil, err
		}
		isLargeEqualResult, err := ttype.TryGetBoolFromValueIf(isLargerEqual)
		if err != nil {
			return nil, err
		}

		if !isLargeEqualResult {
			maxVal = nowVal
		}
	}

	return maxVal.ToInterface()
}

//	Min return the minimum of params with interface mode
func Min(params ...interface{}) (result interface{}, err error) {
	if len(params) <= 0 {
		return nil, nil
	}

	minVal, err := ttype.LoadValueIfFromInterface(params[0])
	if err != nil {
		return nil, err
	}

	//	compare and get minVal
	for _, item := range params {
		nowVal, err := ttype.LoadValueIfFromInterface(item)
		if err != nil {
			return nil, err
		}

		isLessEqual, err := ttype.OpLessEqual.Operate(minVal, nowVal)
		if err != nil {
			return nil, err
		}
		isLessEqualResult, err := ttype.TryGetBoolFromValueIf(isLessEqual)
		if err != nil {
			return nil, err
		}

		if !isLessEqualResult {
			minVal = nowVal
		}
	}

	return minVal.ToInterface()
}