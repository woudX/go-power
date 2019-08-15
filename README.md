Go-Power
===

A lightweight tools library for golang developer, which provider lots of useful tools to make coding easier, faster and feel more comfortable

It contains :
- **container** : provide a series of method to make containers operator more easily, for example :slice, map, etc.
- **convert** : provide common used type & other convert method
- **mathex** : provide common used math/calculate functions
- **powerr** : provider a simple error library which can save dict and stack
- **reflector** : provider common used reflect method
- **structure** : provider kinds of interesting data struct, for example: workcluster, parammap, etc.
- **ttype** : provide reflection based internal type system to process any data

## Goal

- Container
    - [x] Slice support
    - [x] Map support
- Convert
    - [x] ToInt、Int64、Bool、Float64、String
    - [x] ToInterfaceSlice
- MathEx
    - [x] Compare、Sum - Bool/IntSeries/FloatSeries/String support
    - [ ] More type support
- PowErr
    - [x] StoreKV support
    - [x] Stack support
- Reflector
    - [x] GetFunctionName、SetVal support
- Structure
    - [x] ParamMap support
    - [x] WorkCluster support
    - [ ] Set support

## Feature

### Container

GoPower provide a series of c-like methods to operate slice/map

- slice : support Find/Remove series method
- map : support TryGet series method

```go

//  Slice
pos, err := container.FindInSlice(sliceObj, 3)
result, err := container.RemoveFromSlice(sliceObj, 3)

//  Map
err := container.MapTryGet(mapObj, "key", &outVal)

```

### Convert

Golang type convert is an annoying problem, hence GoPower provide a series of common used methods to help develop do converter

```go
intVal, err := convert.ToInt(srcVal)
int64Val, err := convert.ToInt64(srcVal)
float64Val, err := convert.ToFloat64(srcVal)
ifSlice, err := convert.ToInterfaceSlice([]int{1,2,3,4,5})
```

### MathEx

GoPower provide extend math library to do some excited operator

- mathex.Compare : compare any kind of two var!(only for basic type now)
- mathex.Sum : sum any kind of var

```go
//  Compare
result, err := mathex.Compare(3.14, 2)
result, err := mathex.Compare(3.14, "4.233")
result, err := mathex.Compare(false, 33.1)

//  Sum
result, err := mathex.Sum(1,2,33.33,4,5,"3.22", true)
```

### PowErr

GoPower provide a lite extend error library named powerr, it support :

- StoreKV : save k-v into err, all KVs will show when print error
- StoreStack : save err stack trace if needed

```go
err := powerr.New("this is a error").StoreKV("int_key", 3).StoreKV("float_key",3.14).StoreStack()
```

### Reflector

GoPower internal reflect assist lib, but you can use it if require

- GetFunctionName : return function handler name
- SetVal : set a value to ref-val

```go
//  GetFuncName
funcName := reflector.GetFunctionName(reflector.GetFunctionName)

//  SetVal
err := reflector.SetVal(3.14, &targetObj)
```

### structure

GoPower contains some useful structure for developer to save time and make coding comfortable

- ParamMap : a map[string]interface{} based structure, which provide Set/Get/TryGet method

```go
//  ParamMap
pm := parammap.NewParamMap()
pm.Set("key1", 1)
pm.Set("key2", 1.123)
pm.Set("key3", "string")

intVal, err := pm.GetInt("key1")
stringVal, err := pm.GetString("key3")

var floatVal float
err := pm.TryGet("key2", &floatVal)

```

- WorkCluster : a highly packaged goroutine work pool for batch processing

```go
//  WorkCluster - get sum of (1..10000) * 2
wc := workcluster.NewWorkCluster()
wc.StartR(context.TODO(), func(ctx context.Context, in int) (out int) {
	return in * 2
}})

go func() {
	for idx := 1; idx <= 10000; idx++ {
		wc.Push(&WCInputData{X:idx})
	}
    wc.PushDone()
}()

sum := 0
for outValItem := range wc.PopChan() {
	outVal, ok := outValItem.(int)
	sum += outVal
}
```