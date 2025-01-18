package callback

import (
	"sync"
	"time"
)

type DataReporter struct {
	DataChannel chan string
	wg          sync.WaitGroup
}

func NewDataReporter() *DataReporter {
	return &DataReporter{
		DataChannel: make(chan string),
	}
}

func (dr *DataReporter) Start(callback ReportCallback) {
	dr.wg.Add(1)
	go func() {
		defer dr.wg.Done()
		for data := range dr.DataChannel {
			// 模拟上报数据的操作，这里只是简单打印
			if callback != nil {
				callback.Callback(data)
			}
			// 可以替换为实际的网络请求等上报操作
			time.Sleep(1 * time.Second)
		}
	}()
}

func (dr *DataReporter) Stop() {
	close(dr.DataChannel)
	dr.wg.Wait()
}

func (dr *DataReporter) SendData(data string) {
	dr.DataChannel <- data
}
