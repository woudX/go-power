package container

import (
	"gopower/constant"
	"gopower/ttype"
	"math"
)

// GoPower provide a series of c-like method to operate slice, they are realised by reflect
// and non-reflect version (not always):
//
// Find	: find object from slice, including Find/FindLast
// Remove : remove object from slice, including Remove/RemoveFirst/RemoveIf
// Replace : replace slice object to another object Replace/ReplaceIf
// Check : check if slice object match condition - AllOf/AnyOf
//
// All non-reflect version support common used type, including: string, int, int64, float32, float64 and func-mode

//=========================================SLice Find (Reflect Version)=======================================//
//	Internal realization for find slice
func _findInSlice(sliceObj []interface{}, startPos int, endPos int, findObj interface{}) (pos int, err error) {
	pos = -1

	//	Convert findObj to ValueIf
	tmpFindObjValIf, err := ttype.LoadValueIfFromInterface(findObj)
	if err != nil {
		return -1, err
	}

	//	Traversal each obj in range and convert to ValueIf, use operator to check equal
	for idx := startPos; idx < endPos; idx++ {
		valIf, err := ttype.LoadValueIfFromInterface(sliceObj[idx])
		if err != nil {
			return -1, err
		}

		//	If can't operate just skip, because can't operate means not equal
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

//	Internal realization for reverse find slice
func _findInSliceReverse(sliceObj []interface{}, startPos int, endPos int, findObj interface{}) (pos int, err error) {
	pos = -1

	//	Convert findObj to ValueIf
	tmpFindObjValIf, err := ttype.LoadValueIfFromInterface(findObj)
	if err != nil {
		return -1, err
	}

	//	Reverse traversal each obj in range and convert to ValueIf, use operator to check equal
	for idx := endPos - 1; idx >= startPos; idx-- {
		valIf, err := ttype.LoadValueIfFromInterface(sliceObj[idx])
		if err != nil {
			return -1, err
		}

		//	If can't operate just skip, because can't operate means not equal
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

//	Find object in slice and return first pos, return -1 if not find
func FindInSlice(sliceObj []interface{}, findObj interface{}) (pos int, err error) {
	return _findInSlice(sliceObj, 0, len(sliceObj), findObj)
}

//	Find object in slice and return last pos, return -1 if not find
func FindLastInSlice(sliceObj []interface{}, findObj interface{}) (pos int, err error) {
	return _findInSliceReverse(sliceObj, 0, len(sliceObj), findObj)
}
//=========================================SLice Find (Non-Reflect Version)=======================================//
//	Compare function definition
type compareFunc func(interface{}, interface{}) (int, error)

//	If function function definition
type ifFunc func(interface{}) (int, error)

//	Internal realization for findif slice
func _findInSliceIf(sliceObj[] interface{}, startPos int, endPos int, ifFuncHdl ifFunc) (pos int, err error) {
	pos = -1
	//	Traversal each obj in range and convert to ValueIf, use operator to check equal
	for idx := startPos; idx < endPos; idx++ {
		find, err := ifFuncHdl(sliceObj[idx])
		if err != nil {
			return pos, err
		}

		if find != 0 {
			pos = idx
			break
		}
	}

	return pos, err
}

//	Internal realization for find slice, provide user define compare function
func _findInSliceCmp(sliceObj []interface{}, startPos int, endPos int, findObj interface{}, cmpFuncHdl compareFunc) (pos int, err error) {
	pos = -1
	//	Traversal each obj in range and convert to ValueIf, use operator to check equal
	for idx := startPos; idx < endPos; idx++ {
		cmpResult, err := cmpFuncHdl(findObj, sliceObj[idx])
		if err != nil {
			return pos, err
		}

		if cmpResult == 0 {
			pos = idx
			break
		}
	}

	return pos, err
}

//	Internal realization for reverse find slice, provide user define compare function
func _findInSliceReverseCmp(sliceObj []interface{}, startPos int, endPos int, findObj interface{}, cmpFuncHdl compareFunc) (pos int, err error) {
	pos = -1
	//	Traversal each obj in range and convert to ValueIf, use operator to check equal
	for idx := startPos; idx < endPos; idx++ {
		cmpResult, err := cmpFuncHdl(findObj, sliceObj[idx])
		if err != nil {
			return pos, err
		}

		if cmpResult == 0 {
			pos = idx
			break
		}
	}

	return pos, err
}

//	Find object in slice with compare function and return first pos, return -1 if not find
func FindInSliceCmp(sliceObj []interface{}, findObj interface{}, cmpFuncHdl compareFunc) (pos int, err error) {
	return _findInSliceCmp(sliceObj, 0, len(sliceObj), findObj, cmpFuncHdl)
}

//	Find object in slice with compare function and return last pos, return -1 if not find
func FindLastInSliceCmp(sliceObj []interface{}, findObj interface{}, cmpFuncHdl compareFunc) (pos int, err error) {
	return _findInSliceReverseCmp(sliceObj, 0, len(sliceObj), findObj, cmpFuncHdl)
}

//	Find object in slice with if function and return first match pos, return -1 if not find
func FindInSliceIf(sliceObj []interface{}, findIfHdl ifFunc) (pos int, err error) {
	return _findInSliceIf(sliceObj, 0, len(sliceObj), findIfHdl)
}

//	Find string in slice and return pos, return -1 if not find
func FindInSliceString(sliceObj []string, findObj string) (pos int, err error) {
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
func FindInSliceInt64(sliceObj []int64, findObj int64) (pos int, err error) {
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
func FindInSliceInt(sliceObj []int, findObj int) (pos int, err error) {
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
func FindInSliceFloat64(sliceObj []float64, findObj float64) (pos int, err error) {
	pos = -1

	for idx, item := range sliceObj {
		if math.Abs(item-findObj) < constant.EPSILON {
			pos = idx
			break
		}
	}

	return pos, nil
}

//	Find float32 in slice and return pos, return -1 if not find
func FindInSliceFloat32(sliceObj []float32, findObj float32) (pos int, err error) {
	pos = -1

	for idx, item := range sliceObj {
		if math.Abs(float64(item-findObj)) < constant.EPSILON {
			pos = idx
			break
		}
	}

	return pos, nil
}

//=========================================SLice Remove (Reflect Version)=======================================//
//	Remove first remObj from slice
func RemoveFirstSlice(sliceObj []interface{}, remObj interface{}) (result []interface{}, err error) {
	pos, err := _findInSlice(sliceObj, 0, len(sliceObj), remObj)
	if err != nil {
		return sliceObj, err
	}

	// return sliceObj if not found
	if pos == -1 {
		return sliceObj, nil
	}

	result = append(result, sliceObj[:pos]...)
	return append(result, sliceObj[pos+1:]...), nil
}

//	Remove all remObj from slice
func RemoveFromSlice(sliceObj []interface{}, remObj interface{}) (result []interface{}, err error) {
	var lastIdx = -1
	newIdx, err := _findInSlice(sliceObj, lastIdx+1, len(sliceObj), remObj)

	for ; newIdx != -1; newIdx, err = _findInSlice(sliceObj, lastIdx+1, len(sliceObj), remObj) {
		if err != nil {
			return sliceObj, err
		}

		result = append(result, sliceObj[lastIdx+1:newIdx]...)
		lastIdx = newIdx
	}

	//	lastIdx equal -1 means not contain
	if lastIdx == -1 {
		result = sliceObj
	} else {
		result = append(result, sliceObj[lastIdx+1:]...)
	}

	return result, nil
}
//=========================================SLice Remove (Non-Reflect Version)=======================================//\
//	Remove all remObj from slice
func RemoveFromSliceCmp(sliceObj []interface{}, remObj interface{}, cmpFuncHdl compareFunc) (result []interface{}, err error) {
	var lastIdx = -1
	newIdx, err := _findInSliceCmp(sliceObj, lastIdx+1, len(sliceObj), remObj, cmpFuncHdl)

	for ; newIdx != -1; newIdx, err = _findInSliceCmp(sliceObj, lastIdx+1, len(sliceObj), remObj, cmpFuncHdl) {
		if err != nil {
			return sliceObj, err
		}

		result = append(result, sliceObj[lastIdx+1:newIdx]...)
		lastIdx = newIdx
	}

	//	lastIdx equal -1 means not contain
	if lastIdx == -1 {
		result = sliceObj
	} else {
		result = append(result, sliceObj[lastIdx+1:]...)
	}

	return result, nil
}

//	Remove all matched object from slice
func RemoveFromSliceIf(sliceObj []interface{}, ifFuncHdl ifFunc) (result []interface{}, err error) {
	var lastIdx = -1
	newIdx, err := _findInSliceIf(sliceObj, lastIdx+1, len(sliceObj), ifFuncHdl)

	for ; newIdx != -1; newIdx, err = _findInSliceIf(sliceObj, lastIdx+1, len(sliceObj), ifFuncHdl) {
		if err != nil {
			return sliceObj, err
		}

		result = append(result, sliceObj[lastIdx+1:newIdx]...)
		lastIdx = newIdx
	}

	//	lastIdx equal -1 means not contain
	if lastIdx == -1 {
		result = sliceObj
	} else {
		result = append(result, sliceObj[lastIdx+1:]...)
	}

	return result, nil
}