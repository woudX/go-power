# Structure 

## WorkCluster

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
