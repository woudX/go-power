# Examples

## container

```go
//  Some data for prepare
complexSlice := []interface{}{-3,544, true, 22, "string-A", 123, 22, "str-B", 3.1, -23.4, 22, 3.111}
stringSlice := []string{"string", "int", "php", "bool", "golang", "php"}
intSlice := []int64{2,4,6,9,10,11}

//  Find data pos in slice
pos, err := container.Find(interfaceSlice..., 22) // pos=4
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

## convert

```go
//  Try converter any data to all kinds of data

srcVal := "1541123"
intVal, err := convert.ToInt(srcVal)
int64Val, err := convert.ToInt64(srcVal)
float64Val, err := convert.ToFloat64(srcVal)
boolVal, err := convert.ToBool(srcVal)
stringVal, err := convert.ToString(srcVal)

//  Try converter slice to interface slice
ifSlice, err := convert.ToInterfaceSlice([]int{1,2,3,4,5})
ifSlice, err := convert.ToInterfaceSlice([]string{"make", "golang", "powerful"})

```

## mathex

GoPower provide more powerful math method, which can support any value do compare/max/min/sum ... operator 

```go
//  Compare to interface value
cmpResult, err := mathex.Compare(3.14, -5)

//  Sum interface value
sumVal, err := mathex.Sum(3.14, -5, 20.3, 55773)

```

## ttype

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


## Structure 

### WorkCluster

```go
    //	Create a work cluster with multi worker
	wc := workcluster.NewDefaultWorkCluster()
	wc.Start(context.TODO(), func(ctx context.Context, inputData interface{}) (outputData interface{}) {
		fmt.Println("process ", inputData)
		return inputData
	})


	//	Producer gen 10000 nums
	go func() {
		for idx := 1; idx < 10000; idx++ {
			wc.Push(idx)
		}
		wc.PushDone()
	}()

	//	Consumer-A use Pop consumer data
	go func() {
	LA:
		for {
			data, status := wc.Pop()
			switch status {
			case workcluster.PopStatusOk:
				fmt.Println("A finished process : ", data)
			case workcluster.PopStatusClosed:
				fmt.Println("channel closed!")
				break LA
			case workcluster.PopStatusTimeOut:
				time.Sleep(time.Millisecond)
			}
		}
	}()

	//	Consumer-B use PopChan consumer data
	go func() {
	LB:
		for {
			select {
			case data, ok := <-wc.PopChan():
				if !ok {
					fmt.Println("channel closed!")
					break LB
				}
				fmt.Println("B finished process : ", data)

			case <-time.After(time.Millisecond * 50):
				time.Sleep(time.Millisecond)
			}
		}
	}()
```
