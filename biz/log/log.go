package log

import (
	"fmt"
	"time"
)

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
