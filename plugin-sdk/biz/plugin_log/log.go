package plugin_log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

func init() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "log/app.log", // 日志文件名称
		MaxSize:    10,            // 每个日志文件最大大小（MB）
		MaxBackups: 30,            // 保留的旧日志文件最大数量
		MaxAge:     7,             // 保留旧日志文件的最大天数
		Compress:   true,          // 是否压缩旧日志文件
	}
	logrus.SetOutput(lumberjackLogger)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          false,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		DisableQuote:           true,
	})
}

var LogList = []string{}

func Info(message string, a ...any) {
	M := buildLogMessage(message, a...)
	fmt.Printf(M)
	fmt.Printf("\n")
	LogList = append(LogList, M)
}

func Error(message string, a ...any) {
	M := buildLogMessage(message, a...)
	fmt.Printf(M)
	fmt.Printf("\n")
	LogList = append(LogList, M)
}

func Warn(message string, a ...any) {
	M := buildLogMessage(message, a...)
	fmt.Printf(M)
	fmt.Printf("\n")
	LogList = append(LogList, M)
}

func Debug(message string, a ...any) {
	M := buildLogMessage(message, a...)
	logrus.Error(M)
}

func Fatal(message string, a ...any) {
	M := buildLogMessage(message, a...)
	fmt.Printf(M)
	fmt.Printf("\n")
	LogList = append(LogList, M)
}

func buildLogMessage(message string, a ...any) string {
	now := time.Now()
	// 格式化时间
	formattedTime := now.Format("2006-01-02 15:04:05")
	return formattedTime + " " + fmt.Sprintf(message, a...)
}

func ClearLogList() []string {
	copyOfLog := append(LogList)
	LogList = []string{}
	return copyOfLog
}
