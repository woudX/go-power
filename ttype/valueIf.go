package ttype

type ValueType string

const (
	IntType    ValueType = "int64"
	FloatType  ValueType = "float64"
	StringType ValueType = "string"
	BoolType   ValueType = "bool"
	SliceType  ValueType = "slice"
)

//	A ValueIf contains multi part
//  - Basic Func : Type/Data
//  - Operation Func : Larger/Smaller/.....
//  - Convertor : ToXXX
type ValueIf interface {
}