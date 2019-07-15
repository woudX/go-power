package mathex

import "github.com/woudX/gopower/ttype"

//	Sum return the sum value of params
func Sum(params ...interface{}) (result interface{}, err error) {
	resultValIf, _ := ttype.LoadValueIfFromInterface(int64(0))

	for _, item := range params {
		itemIf, err := ttype.LoadValueIfFromInterface(item)
		if err != nil {
			return nil, err
		}

		if resultValIf, err  = ttype.OpAdd.Operate(resultValIf, itemIf); err != nil {
			return nil, err
		}
	}

	return resultValIf.ToInterface()
}