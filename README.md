Go-Power
===

> Library is under developing now, I desire to use lib after release version published

A lightly tool library for go developer, which provider lots of useful tool to make coding easier, faster and feel more comfortable

It contains :
- container : provide more function to operate containers in golang, for example: slice
- converter : provide common used type & other convert method
- mathex : provide common used math/calculate functions
- ttype : provide reflect method to make interface{} more easy to use

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