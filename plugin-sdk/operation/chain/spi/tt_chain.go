package spi

var TtReportCallbackInstance = &TtReportCallback{}

type TtReportCallback struct {
}

func (*TtReportCallback) Callback(input interface{}) {
}

func (*TtReportCallback) StopCallback() {

}
func (*TtReportCallback) OverWriteTt() {

}
