package container

import (
	"gopower/constant"
	"gopower/ttype"
	"math"
)

//	Find any object in slice and return pos, return -1 if not find
func SliceFind(sliceObj []interface{}, findObj interface{}) (pos int, err error) {
	pos = -1

	//	Convert findObj to ValueIf
	tmpFindObjValIf, err := ttype.LoadValueFromInterface(findObj)
	if err != nil {
		return -1, err
	}

	//	Traversal each obj in slice and convert to ValueIf, use operator to check equal
	for idx, item := range sliceObj {
		valIf, err := ttype.LoadValueFromInterface(item)
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

//	Find string in slice and return pos, return -1 if not find
func SliceFindString(sliceObj []string, findObj string) (pos int, err error) {
	pos = -1

	for idx, item := range sliceObj {
		if item == findObj {
			pos = idx
			break
		}
	}

	return pos, nil
}

//	Find int64 in slice and return pos, return -1 if not find
func SliceFindInt64(sliceObj []int64, findObj int64) (pos int, err error) {
	pos = -1

	for idx, item := range sliceObj {
		if item == findObj {
			pos = idx
			break
		}
	}

	return pos, nil
}

//	Find int in slice and return pos, return -1 if not find
func SliceFindInt(sliceObj []int, findObj int) (pos int, err error) {
	pos = -1

	for idx, item := range sliceObj {
		if item == findObj {
			pos = idx
			break
		}
	}

	return pos, nil
}

//	Find float64 in slice and return pos, return -1 if not find
func SliceFindFloat64(sliceObj []float64, findObj float64) (pos int, err error) {
	pos = -1

	for idx, item := range sliceObj {
		if math.Abs(item - findObj) < constant.EPSILON {
			pos = idx
			break
		}
	}

	return pos, nil
}

//	Find float32 in slice and return pos, return -1 if not find
func SliceFindFloat32(sliceObj []float32, findObj float32) (pos int, err error) {
	pos = -1

	for idx, item := range sliceObj {
		if math.Abs(float64(item - findObj)) < constant.EPSILON {
			pos = idx
			break
		}
	}

	return pos, nil
}