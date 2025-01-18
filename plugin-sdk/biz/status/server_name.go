package status

const (
	CQTT = iota // 宠爱天堂
	KDHS        // 幻世
)

var SERVER_NAME = KDHS

// 是否使同账号的sessionid保持一致
var SESSION_ID_KEEP_TYPE = true
