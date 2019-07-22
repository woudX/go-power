package workcluster

import "context"

//	WorkHandler function define
type workHandlerFunc func(ctx context.Context, inputData interface{}) (outputData interface{})

//	WorkCluster interface
type WorkCluster interface {
	Start(ctx context.Context, workHdl workHandlerFunc)
	StartR(ctx context.Context, workHdl interface{})

	Push(inputDataList ...interface{}) WorkCluster
	PushDone() WorkCluster

	PopChan() chan interface{}
	Pop() (result interface{}, status PopStatus)
	PopT(blockMillisecond int64) (result interface{}, status PopStatus)
	TryPop(referenceOut interface{}) (status PopStatus, err error)
	TryPopT(referenceOut interface{}, blockMillisecond int64) (status PopStatus, err error)

	Wait() WorkCluster
	Stop() WorkCluster
}
