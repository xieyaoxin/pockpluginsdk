package spi

var DungeonInstanceReportCallbackInstance = &DungeonInstanceReportCallback{}

type DungeonInstanceReportCallback struct {
}

func (*DungeonInstanceReportCallback) Callback(input interface{}) {
}

func (*DungeonInstanceReportCallback) StopCallback() {

}
func (*DungeonInstanceReportCallback) OverWriteDungeon() {

}
