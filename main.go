package main

import (
	"context"
	"fmt"
	"github.com/woudX/gopower/container"
	"github.com/woudX/gopower/convert"
	"github.com/woudX/gopower/mathex"
	"github.com/woudX/gopower/structure/workcluster"
	"math/rand"
	"time"
)

var demoInput = []interface{}{-3, 544, true, 22, "string-A", 123, "str-B", 3.1, -23.4, 3.111}
var demoMultiInput = []interface{}{-3, 544, true, 22, "string-A", 123, 22, 22, "str-B", 3.1, -23.4, 22, 3.111}
var demoStringInput = []string{"aaa", "string-B", "string-A"}
var demoIntInput = []int{2, 4, 6, 9, 10, 11}

type demoStruct struct {
	fieldA int     `json:"field_a"`
	fieldB *string `json:"field_b"`
	fieldC int64   `json:"field_c"`
}

func (ds *demoStruct) String() string {
	return fmt.Sprintf("field_a=%v, field_b=%v, field_c=%v", ds.fieldA, ds.fieldB, ds.fieldC)
}

func demoFunc(ctx context.Context, reqData *demoStruct) interface{} {
	return 1
}

func main() {
	wc := workcluster.NewDefaultWorkCluster()
	ctx := context.WithValue(context.TODO(), "traceid", "12345")
	wc.StartR(ctx, func(ctx context.Context, intVal *int) *demoStruct {
		fmt.Println("func_val=", intVal)
		fmt.Println("trace_id=", ctx.Value("traceid"))

		return &demoStruct{
			fieldA: *intVal,
			fieldC: time.Now().Unix(),
		}
	})

	go func() {
		for idx := 0; idx < 20; idx++ {
			wc.Push(idx)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}

		wc.PushDone()
	}()

	for val := range wc.PopChan() {
		fmt.Println("rcv %+v", val.(*demoStruct))
	}

	if wc.Err != nil {
		fmt.Println(wc.Err.Error())
	}

	return
	pos, err := container.FindInSlice(convert.MustToInterfaceSlice(demoStringInput), "string-A")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(pos)

	result, err := container.FindInSliceIf(convert.MustToInterfaceSlice(demoIntInput), func(val interface{}) (result int, err error) {
		return convert.ToInt(val.(int)%2 == 1)
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)

	result2, err := container.RemoveFromSliceIf(convert.MustToInterfaceSlice(demoIntInput), func(val interface{}) (result int, err error) {
		return convert.ToInt(val.(int)%2 == 1)
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result2)

	//	Test-Compare
	//	TODO: string(float) vs int may cause problem
	fmt.Println(mathex.Compare(3, 2.9999))
	fmt.Println(mathex.Compare(3, 2.9999999999999))
	fmt.Println(mathex.Compare(3, true))
	fmt.Println(mathex.Compare(0, true))
	fmt.Println(mathex.Compare("a", "aa"))
	fmt.Println(mathex.Compare("b", "aa"))

	fmt.Println(mathex.Max(3, 2, 1, 3.33, 2, true, "3.415"))
	fmt.Println(mathex.Max("3.415"))

	fmt.Println(mathex.Min(3, 2, 1, 3.33, 2, true, "3.415"))
	fmt.Println(mathex.Min("3.415"))

	fmt.Println(mathex.Sum(3, 2, 1, 3.33, 2, "3.415"))

	time.Sleep(time.Second * 1000)
}
