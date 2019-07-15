package ttype

type ValueType string

const (
	TypeInt64   ValueType = "int64"
	TypeFloat64 ValueType = "float64"
	TypeString  ValueType = "string"
	TypeBool    ValueType = "bool"
	TypeSlice   ValueType = "slice"
)

//	A ValueIf contains multi part
//  - Basic Func : Type/Data
//  - Operation Func : Larger/Smaller/.....
//  - Convertor : ToXXX
type ValueIf interface {
	//	Basic
	Type() ValueType
	Priority() int
	Value() (interface{}, error)

	//	Operations
	Equal(rhs ValueIf) (ValueIf, error)
	Larger(rhs ValueIf) (ValueIf, error)
	LargerEqual(rhs ValueIf) (ValueIf, error)
	Less(rhs ValueIf) (ValueIf, error)
	LessEqual(rhs ValueIf) (ValueIf, error)

	//	Converter
	ToInt64() (*ValueInt64, error)
	ToFloat64() (*ValueFloat64, error)
	ToBool() (*ValueBool, error)
	ToString() (*ValueString, error)
	ToInterface() (interface{}, error)
}
