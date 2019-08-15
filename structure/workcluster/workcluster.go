package workcluster

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"

	"github.com/woudX/gopower/powerr"
	"github.com/woudX/gopower/reflector"
)

type PopStatus int

const (
	PopStatusOk PopStatus = iota
	PopStatusTimeOut
	PopStatusClosed
	PopStatusError
)

var DefaultPopBlockMillisecond int64 = 50

//	workCluster recv a task and create multi worker goroutine to process data
//  - Select worker nums
//  - Easy create and post data
//  - Goroutine safe and return data easily, which kind of types you like
type workCluster struct {
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
func NewWorkCluster() WorkCluster {
	return NewWorkClusterCustom(20, 100)
}

//	Create a work cluster
func NewWorkClusterCustom(workerCount int, chanCache int) WorkCluster {
	cluster := &workCluster{
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

//	Error return workCluster running error info
func (wc *workCluster) Error() error {
	return wc.Err
}

//	StartR use reflect package to fill workHdl params and call function handler method
//	accept a context.context and a custom function handler as input. It should be noted that
//	[function params] type need equal to [input channel] type, for example:
//
//	type RequestStruct struct {...}
//	wc.StartR(ctx, func(ctx context.context, requestData *RequestStruct) (bool))
//
//	This define need input channel data type *RequestStruct and will return bool type as result
func (wc *workCluster) StartR(ctx context.Context, workHdl interface{}) WorkCluster {
	//	set wait group
	wc.waiter = &sync.WaitGroup{}
	wc.waiter.Add(wc.workerCount)

	//	start goroutine
	for idx := 0; idx < wc.workerCount; idx++ {
		go func(gid int, inputChan chan interface{}, outputChan chan interface{}, ctrlChan chan int) {
			defer func() {
				if r := recover(); r != nil {
					wc.Err = powerr.New("workCluster.StartR panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
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
		defer func() {
			if r := recover(); r != nil {
				wc.Err = powerr.New("workCluster.StartR panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
			}

			wc.waiter.Done()
		}()

		wc.waiter.Wait()
		close(wc.outputChan)
	}()

	return wc
}

//	Start worker goroutines and process input data
//	method accept a context.context and workHandlerFunc as input, for example:
//
//	wc.Start(ctx, func(context.context, interface{}) interface{})
//
//	if you want use custom type as input and output, we suggest using StartR() method
func (wc *workCluster) Start(ctx context.Context, workHdl workHandlerFunc) WorkCluster {
	//	set wait group
	wc.waiter = &sync.WaitGroup{}
	wc.waiter.Add(wc.workerCount)

	//	start goroutine
	for idx := 0; idx < wc.workerCount; idx++ {
		go func(gid int, inputChan chan interface{}, outputChan chan interface{}, ctrlChan chan int) {
			defer func() {
				if r := recover(); r != nil {
					wc.Err = powerr.New("workCluster.Start panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
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

	return wc
}

//	Push will send inputData to channel, and then processed by worker
func (wc *workCluster) Push(inputDataList ...interface{}) WorkCluster {
	defer func() {
		if r := recover(); r != nil {
			wc.Err = powerr.New("workCluster.Push panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
		}
	}()

	for _, data := range inputDataList {
		randIdx := rand.Intn(wc.workerCount)
		wc.inputChanList[randIdx] <- data
	}

	return wc
}

//	PushDone close all input channel and finish push
func (wc *workCluster) PushDone() WorkCluster {
	defer func() {
		if r := recover(); r != nil {
			wc.Err = powerr.New("workCluster.PushDone panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
		}
	}()

	for idx := range wc.inputChanList {
		close(wc.inputChanList[idx])
	}

	return wc
}

//	Wait until all worker finished
//	This method may cause blocking
func (wc *workCluster) Wait() WorkCluster {
	wc.waiter.Wait()
	return wc
}

//	PopT return a result from output chan with input blockMillisecond
func (wc *workCluster) PopT(blockMillisecond int64) (interface{}, PopStatus) {
	defer func() {
		if r := recover(); r != nil {
			wc.Err = powerr.New("workCluster.PopT panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
		}
	}()

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
func (wc *workCluster) Pop() (interface{}, PopStatus) {
	return wc.PopT(DefaultPopBlockMillisecond)
}

//	TryPopT return result with custom data type with reflector.SetVal in blockMillisecond
func (wc *workCluster) TryPopT(blockMillisecond int64, referenceOut interface{}) (status PopStatus, err error) {

	//	Pope data from channel with block time
	val, popStatus := wc.PopT(blockMillisecond)

	switch popStatus {
	case PopStatusOk:
		if err = reflector.SetVal(val, referenceOut); err != nil {
			return PopStatusError, err
		}

		return popStatus, nil
	default:
		return popStatus, nil
	}
}

//	TryPop insist return result with custom data type with reflector.SetVal
func (wc *workCluster) TryPop(referenceOut interface{}) (popStatus PopStatus, err error) {
	return wc.TryPopT(DefaultPopBlockMillisecond, referenceOut)
}

//	PopChan return the output channel
func (wc *workCluster) PopChan() chan interface{} {
	return wc.outputChan
}

//	Stop the work cluster, all remain data in input chan will lose
func (wc *workCluster) Stop() WorkCluster {
	defer func() {
		if r := recover(); r != nil {
			wc.Err = powerr.New("workCluster.Stop panic occur").StoreKV("panic_info", fmt.Sprintf("%v", r))
		}
	}()

	for idx := range wc.ctrlChanList {
		close(wc.ctrlChanList[idx])
	}

	return wc
}