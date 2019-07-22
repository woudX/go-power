package set

type SetIf interface {
	Insert(interface{})
	Remove(interface{})
	Contain(interface{})
	Clear()
}
