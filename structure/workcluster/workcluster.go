package workcluster

import (
	"context"
	"fmt"
	"github.com/woudX/gopower/powerr"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

type PopStatus int

const (
	PopStatusOk PopStatus = iota
	PopStatusTimeOut
	PopStatusClosed
)

var DefaultPopBlockMillisecond int64 = 50

//	WorkHandler function define
type workHandlerFunc func(ctx context.Context, inputData interface{}) (outputData interface{})

//	WorkCluster recv a task and create multi worker goroutine to process data
//  - Select worker nums
//  - Easy create and post data
//  - Goroutine safe and return data easily, which kind of types you like
type WorkCluster struct {
	//	worker counter
	workerCount int

	//	wait lock
	waiter *sync.WaitGroup

	//	communication channel
	inputChanList []chan interface{}
	outputChan    chan interface{}

	//	control channel
	ctrlChanList []chan int

	//	error
	Err error
}

//	Create a default param work cluster
func NewDefaultWorkCluster() *WorkCluster {
	return NewWorkCluster(20, 100)
}

//	Create a work cluster
func NewWorkCluster(workerCount int, chanCache int) *WorkCluster {
	cluster := &WorkCluster{
		workerCount: workerCount,
	}

	cluster.inputChanList = make([]chan interface{}, 0, chanCache)
	cluster.outputChan = make(chan interface{}, chanCache*workerCount)
	for idx := 0; idx < cluster.workerCount; idx++ {
		cluster.inputChanList = append(cluster.inputChanList, make(chan interface{}, chanCache))
		cluster.ctrlChanList = append(cluster.ctrlChanList, make(chan int, chanCache))
	}

	return cluster
}

//	StartR use reflect package to fill workHdl params and call function handler method
//	accept a context.context and a custom function handler as input. It should be noted that
//	[function params] type need equal to [input channel] type, for example:
//
//	type RequestStruct struct {...}
//	wc.StartR(ctx, func(ctx context.context, requestData *RequestStruct) (bool))
//
//	This define need input channel data type *RequestStruct and will return bool type as result
func (wc *WorkCluster) StartR(ctx context.Context, workHdl interface{}) {
	//	set wait group
	wc.waiter = &sync.WaitGroup{}
	wc.waiter.Add(wc.workerCount)

	//	start goroutine
	for idx := 0; idx < wc.workerCount; idx++ {
		go func(gid int, inputChan chan interface{}, outputChan chan interface{}, ctrlChan chan int) {
			defer func() {
				if r := recover(); r != nil {
					wc.Err = powerr.New("workCluster panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
				}

				wc.waiter.Done()
			}()

		exitLoop:
			for {
				select {
				case <-ctrlChan:
					break exitLoop

				case inputData, ok := <-inputChan:
					if !ok {
						break exitLoop
					}

					//	use reflect method to call func and return data
					funcVal := reflect.ValueOf(workHdl)
					returnValList := funcVal.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(inputData)})
					if len(returnValList) > 0 {
						outputChan <- returnValList[0].Interface()
					}

					//	when process can't read data from inputChan, it sleep 1 millisecond to return thread controller
				case <-time.After(time.Millisecond * 50):
					time.Sleep(time.Millisecond)
				}
			}

		}(idx, wc.inputChanList[idx], wc.outputChan, wc.ctrlChanList[idx])
	}

	//	start goroutine for detect worker finished and close output channel
	go func() {
		wc.waiter.Wait()
		close(wc.outputChan)
	}()
}

//	Start worker goroutines and process input data
//	method accept a context.context and workHandlerFunc as input, for example:
//
//	wc.Start(ctx, func(context.context, interface{}) interface{})
//
//	if you want use custom type as input and output, we suggest using StartR() method
func (wc *WorkCluster) Start(ctx context.Context, workHdl workHandlerFunc) {
	//	set wait group
	wc.waiter = &sync.WaitGroup{}
	wc.waiter.Add(wc.workerCount)

	//	start goroutine
	for idx := 0; idx < wc.workerCount; idx++ {
		go func(gid int, inputChan chan interface{}, outputChan chan interface{}, ctrlChan chan int) {
			defer func() {
				if r := recover(); r != nil {
					wc.Err = powerr.New("workCluster panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
				}

				wc.waiter.Done()
			}()

		exitLoop:
			for {
				select {
				case <-ctrlChan:
					break exitLoop

				case inputData, ok := <-inputChan:
					if !ok {
						break exitLoop
					}
					outputChan <- workHdl(ctx, inputData)

					//	when process can't read data from inputChan, it sleep 1 millisecond to return thread controller
				case <-time.After(time.Millisecond * 50):
					time.Sleep(time.Millisecond)
				}
			}

		}(idx, wc.inputChanList[idx], wc.outputChan, wc.ctrlChanList[idx])
	}

	//	start goroutine for detect worker finished and close output channel
	go func() {
		wc.waiter.Wait()
		close(wc.outputChan)
	}()
}

//	Push will send inputData to channel, and then processed by worker
func (wc *WorkCluster) Push(inputDataList ...interface{}) *WorkCluster {
	for _, data := range inputDataList {
		randIdx := rand.Intn(wc.workerCount)
		wc.inputChanList[randIdx] <- data
	}

	return wc
}

//	PushDone close all input channel and finish push
func (wc *WorkCluster) PushDone() *WorkCluster {
	for idx := range wc.inputChanList {
		close(wc.inputChanList[idx])
	}

	return wc
}

//	Wait until all worker finished
//	This method may cause blocking
func (wc *WorkCluster) Wait() *WorkCluster {
	wc.waiter.Wait()
	return wc
}

//	PopT return a result from output chan with input blockMillisecond
func (wc *WorkCluster) PopT(blockMillisecond int64) (interface{}, PopStatus) {
	select {
	case <-time.After(time.Duration(blockMillisecond) * time.Millisecond):
		return nil, PopStatusTimeOut
	case val, ok := <-wc.outputChan:
		if !ok {
			return nil, PopStatusClosed
		} else {
			return val, PopStatusOk
		}
	}
}

//	Pop return a result from output chan with 50ms block time
func (wc *WorkCluster) Pop() (interface{}, PopStatus) {
	return wc.PopT(DefaultPopBlockMillisecond)
}

//	PopChan return the output channel
func (wc *WorkCluster) PopChan() chan interface{} {
	return wc.outputChan
}

//	Stop the work cluster, all remain data in input chan will lose
func (wc *WorkCluster) Stop() {
	for idx := range wc.ctrlChanList {
		close(wc.ctrlChanList[idx])
	}
}
