package container

import "gopower/ttype"

//	Find object from targets and return position, return -1 if not exist
func Find(findObj interface{}, targets ...interface{}) (pos int, err error) {
	pos = -1

	//	Convert findObj to ValueIf
	tmpFindObjValIf, err := ttype.LoadValueIfFromInterface(findObj)
	if err != nil {
		return -1, err
	}

	//	Traversal each obj in targets and convert to ValueIf, use operator to check equal
	for idx, item := range targets {
		valIf, err := ttype.LoadValueIfFromInterface(item)
		if err != nil {
			return -1, err
		}

		//	If can't operate just skip, becuase can't operate means not equal
		result, err := ttype.OpEqual.Operate(tmpFindObjValIf, valIf)
		if err != nil {
			continue
		}

		finded, err := ttype.TryGetBoolFromValueIf(result)
		if err != nil {
			return -1, err
		}

		if finded {
			pos = idx
			break
		}
	}

	return pos, err
}