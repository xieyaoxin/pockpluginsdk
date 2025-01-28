package biz_callback

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
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
			} else {
				plugin_log.Info("未配置战斗结束回调")
			}
			// 可以替换为实际的网络请求等上报操作
			time.Sleep(1 * time.Second)
		}
	}()
}

func (dr *DataReporter) Stop(callback ReportCallback) {
	close(dr.DataChannel)
	if callback != nil {
		callback.StopCallback()
	} else {
		plugin_log.Info("未配置挂机结束回调")
	}
	dr.wg.Wait()
}

func (dr *DataReporter) SendData(data string) {
	dr.DataChannel <- data
}
