Go-Power
===

> Library is **under developing now**,it's better to use after release version published

A lightweight tools library for golang developer, which provider lots of useful tools to make coding easier, faster and feel more comfortable

It contains :
- **container** : provide a series of method to make containers operator more easily, such as slice, etc.
- **convert** : provide common used type & other convert method
- **mathex** : provide common used math/calculate functions
- **ttype** : provide reflection based internal type system to process interface
- **powerr** : provider a simple error library which can save dict and stack
- **reflector** : provider common used reflect method

## Examples

### container

GoPower provide a series of c-like methods to operate slice, include remove, find, etc. They are all named with SliceXXX, for example:

```go
//  Some data for prepare
complexSlice := []interface{}{-3,544, true, 22, "string-A", 123, 22, "str-B", 3.1, -23.4, 22, 3.111}
stringSlice := []string{"string", "int", "php", "bool", "golang", "php"}
intSlice := []int64{2,4,6,9,10,11}

//  GoPower provide Find, FindInSlice, FindInSliceXXX and FindIf method

//  Find data pos in slice
pos, err := container.Find(complexSlice..., 22) // pos=4
pos, err := container.FindInSlice(complexSlice, 3.1) // pos=7
pos, err := container.FindInSliceString(stringSlice, "bool") // pos=2
pos, err := container.FindInSliceInt64(intSlice, 444) // pos=-1

pos, err := container.FindLastInSlice(complexSlice, 22) // pos=10

pos, err := container.FindInSliceIf(convert.MustToInterfaceSlice(intSlice), func(val interface{}) (result int, err error) {
	return convert.ToInt(val.(int) % 2 == 1)}) // pos=3 

pos, err := container.FindInSliceCmp(...) // use customed functions

//  Remove data from slice
rmSlice, err := container.RemoveFromSlice(complexSlice, 22) // Remove all 22
rmSlice, err := container.RemoveFromSlice(convert.MustToInterfaceSlice(stringSlice), "php") // Remove all "php"

rmSlice, err := container.RemoveFirstFromSlice(complexSlice, 22) // Remove first 22

rmSlice, err := container.RemoveFromSliceIf(convert.MustToInterfaceSlice(intSlice), func(val interface{}) (result int, err error) {
	return convert.ToInt(val.(int) % 2 == 1)}) // Remove all Odd(9, 11)

rmSlice, err := container.RemoveFromSliceCmp(...) // use customed functions

```

### convert

Golang type convert is an annoying problem, so this library provide a series of common used methods to help develop do converter

```go
//  Try converter any data to all kinds of data

srcVal := "1541123"
intVal, err := converter.ToInt(srcVal)
int64Val, err := converter.ToInt64(srcVal)
float64Val, err := converter.ToFloat64(srcVal)
boolVal, err := converter.ToBool(srcVal)
stringVal, err := converter.ToString(srcVal)

//  Try converter slice to interface slice
ifSlice, err := converter.ToInterfaceSlice([]int{1,2,3,4,5})
ifSlice, err := converter.ToInterfaceSlice([]string{"make", "golang", "powerful"})

```

### mathex

GoPower provide more powerful math method, which can support any value do compare/max/min/sum ... operator 

```go
//  Compare to interface value
cmpResult, err := mathex.Compare(3.14, -5)

//  Get Max/Min interface value
maxVal, err := mathex.Max(3.14, -5, 20.3, 55773, true)
minVal, err := mathex.Min(3.14, -5, 20.3, 55773, true)

//  Sum interface value
sumVal, err := mathex.Sum(3.14, -5, 20.3, 55773)

```

### ttype

GoPower use reflection to realize an internal type system called ttype. It often used for another part , but it can also be used directly. The most important type is ValueIf, it's a interface which support many useful method, for example:

```go
//  Create ValueIf from interface
float64ValIf, err := ttype.LoadValueIfFromInterface(3.14) 
int64ValIf, err := ttype.LoadValueIfFromInterface(3) 

//  Get ValIf info
float64ValIf.Type() // float64
int64ValIf.Type() // int64 

//  Convert and get value
cvtInt64Val := float64ValIf.ToInt64().Value() // 3.14 --> 3 

//  Do equal operator and get result
resultValIf := ttype.OpEqual.Operate(float64ValIf, int64ValIf)
boolResult, err := ttype.TryGetBoolFromValueIf(resultValIf) // boolResult is bool = false
```
