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
//  If you want find string exist in slice:
pos, err := container.SliceFind([]string{"string", "int", "bool"}, "bool")

//  You can even find object exist in interface slice:
pos, err := container.SliceFind([]interface{}{-3,544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111}, 3.1)

//  If need speed, suggest use non-reflection function:
pos, err := container.SliceFindString([]string{"string", "int", "bool"}, "bool")

```

- converter

```go

//  Try converter any data to all kinds of data
intVal, err := converter.ToInt(interface{})
int64Val, err := converter.ToInt64(interface{})
float64Val, err := converter.ToFloat64(interface{})
boolVal, err := converter.ToBool(interface{})
stringVal, err := converter.ToString(interface{})

```

- mathex

```go
//  Compare to interface value
cmpResult, err := mathex.Compare(3.14, -5)

//  Get Max/Min interface value
maxVal, err := mathex.Max(3.14, -5, 20.3, 55773, true)
minVal, err := mathex.Min(3.14, -5, 20.3, 55773, true)

//  Sum interface value
sumVal, err := mathex.Sum(3.14, -5, 20.3, 55773)
```