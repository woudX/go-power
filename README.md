Go-Power
===

> Library is under developing now, I desire to use lib after release version published

A lightly tool library for go developer, which provider lots of useful tool to make coding easier, faster and feel more comfortable

It contains :
- container : provide a series of method to make containers operator easy, such as slice, etc.
- converter : provide common used type & other convert method
- mathex : provide common used math/calculate functions
- ttype : provide reflection based internal type system to process interface
- powerr : provider a simple error library which can save dict and stack
- reflector : provider common used reflect method

Example:

- container

```go

// GoPower provide a series of methods like c++ to operate slice, include remove, find ...
// They are all named with SliceXXX, for example:

//  Find data pos in slice
pos, err := container.SliceFind([]interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111}, 3.1)

//  Find data pos in slice (non-reflect version)
pos, err := container.SliceFindString([]string{"string", "int", "bool"}, "bool")
pos, err := container.SliceFindInt64([]int64{111,222,333}, 222)

//  Remove data from slice
rmSlice, err := container.SliceRemove([]interface{"string", "int", "bool"}, "int")

//  Remove data from slice (non-reflect version)
rmSlice, err := container.SliceRemoveString([]string{"string", "int", "bool"}, "int")

```

- converter

```go
//  Golang type convert is an annoying problem, so this library provide a series of common used methods to help
//  develop do converter

//  Try converter any data to all kinds of data
intVal, err := converter.ToInt(interface{})
int64Val, err := converter.ToInt64(interface{})
float64Val, err := converter.ToFloat64(interface{})
boolVal, err := converter.ToBool(interface{})
stringVal, err := converter.ToString(interface{})

```

- mathex

```go
//  GoPower provide more powerful math method, which can support any value do compare/max/min/sum ... operator 

//  Compare to interface value
cmpResult, err := mathex.Compare(3.14, -5)

//  Get Max/Min interface value
maxVal, err := mathex.Max(3.14, -5, 20.3, 55773, true)
minVal, err := mathex.Min(3.14, -5, 20.3, 55773, true)

//  Sum interface value
sumVal, err := mathex.Sum(3.14, -5, 20.3, 55773)

```

- ttype

```go
// GoPower use reflection to realize an internal type system called ttype. It often used for another part
// , but it can also be used directly. Most important type is ValueIf, it's a interface which support many 
// useful method, for example:

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
