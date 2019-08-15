package workcluster

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type WCInputData struct {
	X int
}

type WCOutputData struct {
	XP  int
	err error
}

func TestNewWorkCluster(t *testing.T) {
	wc := NewWorkClusterCustom(10, 100)

	wc.StartR(context.TODO(), func(ctx context.Context, data *WCInputData) (out *WCOutputData) {
		return &WCOutputData{
			XP:  data.X * 2,
			err: nil,
		}
	})

	go func() {
		for idx := 1; idx <= 100; idx++ {
			wc.Push(&WCInputData{X:idx})
		}

		wc.PushDone()
	}()

	sum := 0
	for wcOutItem := range wc.PopChan() {
		outVal, ok := wcOutItem.(*WCOutputData)
		assert.Equal(t, true, ok)
		sum += outVal.XP
	}

	assert.Equal(t, 10100, sum)
}

func TestNewWorkCluster2(t *testing.T) {
	wc := NewWorkClusterCustom(10, 100)

	wc.StartR(context.TODO(), func(ctx context.Context, data *WCInputData) (out *WCOutputData) {
		return &WCOutputData{
			XP:  data.X * 2,
			err: nil,
		}
	})

	go func() {
		for idx := 1; idx <= 100; idx++ {
			wc.Push(&WCInputData{X:idx})
		}

		wc.PushDone()
	}()

	sum := 0
	for {
		wcOutItem, status := wc.Pop()

		if status == PopStatusOk {
			outVal, ok := wcOutItem.(*WCOutputData)
			assert.Equal(t, true, ok)
			sum += outVal.XP
		}

		if status == PopStatusClosed {
			break
		}
	}

	assert.Equal(t, 10100, sum)
}

func TestNewWorkCluster3(t *testing.T) {
	wc := NewWorkClusterCustom(10, 100)

	wc.Start(context.TODO(), func(ctx context.Context, data interface{}) (out interface{}) {

		inData, ok := data.(*WCInputData)
		if !ok {
			return &WCOutputData{
				XP: 0,
				err: errors.New("invalid input data"),
			}
		}

		return &WCOutputData{
			XP:  inData.X * 2,
			err: nil,
		}
	})

	go func() {
		for idx := 1; idx <= 100; idx++ {
			wc.Push(&WCInputData{X:idx})
		}

		wc.PushDone()
	}()

	sum := 0
	for {
		wcOutItem, status := wc.Pop()

		if status == PopStatusOk {
			outVal, ok := wcOutItem.(*WCOutputData)
			assert.Equal(t, true, ok)
			sum += outVal.XP
		}

		if status == PopStatusClosed {
			break
		}
	}

	assert.Equal(t, 10100, sum)
}
