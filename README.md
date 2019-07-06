Go-Power
===

> Library is developing now, we desire to use after release version publish

A useful tool library for go developer, which provider lots of power tool to make coding easier, faster and feel comfortable

It contains :
- container : provide more function to operate containers in golang, for example: slice
- ttype : make interface{} more easy to use, provide Interface to any basic data type
- math : provide most often used math/calculate functions


Example:

- container

```go
//  If you want find string exist in slice:
pos, err := container.SliceFind([]string{"string", "int", "bool"}, "bool")

//  You can even find object exist in interface slice:
pos, err := container.SliceFind([]string{"string", 3, 7.14, -321, 3455}, -321)

//  If need speed, suggest use non-reflection function:
pos, err := container.SliceFindString([]string{"string", "int", "bool"}, "bool")

```

- ttype

```go

```
