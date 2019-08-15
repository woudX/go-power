package workcluster

import "context"

//	WorkHandler function define
type workHandlerFunc func(ctx context.Context, inputData interface{}) (outputData interface{})

//	WorkCluster interface
type WorkCluster interface {
	Error() error

	Start(ctx context.Context, workHdl workHandlerFunc) WorkCluster
	StartR(ctx context.Context, workHdl interface{}) WorkCluster

	Push(inputDataList ...interface{}) WorkCluster
	PushDone() WorkCluster

	PopChan() chan interface{}
	Pop() (result interface{}, status PopStatus)
	PopT(blockMillisecond int64) (result interface{}, status PopStatus)
	TryPop(referenceOut interface{}) (status PopStatus, err error)
	TryPopT(blockMillisecond int64, referenceOut interface{}) (status PopStatus, err error)

	Wait() WorkCluster
	Stop() WorkCluster
}
