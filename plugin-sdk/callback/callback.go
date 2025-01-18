package biz_callback

type ReportCallback interface {
	Callback(data interface{})
	StopCallback()
}
