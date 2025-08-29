package logger

import (
	"fmt"
	"os"
)

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Green  = "\033[32m"
)

type Level int

const (
    INFO Level = iota
    WARN
    ERROR
	SUCCESS
)

func FormatLog(level Level, format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	switch level {
	case INFO:
		return fmt.Sprintf("%si%s %s", Blue, message, Reset)
	case WARN:
		return fmt.Sprintf("%s[WARN]%s %s", Yellow, message, Reset)
	case ERROR:
		return fmt.Sprintf("%s[ERROR]%s %s", Red, Reset, message)
	case SUCCESS:
		return fmt.Sprintf("%s+%s %s", Green, message, Reset)
	default:
		return message
	}
}


func Log(level Level, format string, a ...interface{}) {
	fmt.Println(FormatLog(level, format, a...))
}

func Fatal(format string, a ...interface{}) {
    Log(ERROR, format, a...)
    os.Exit(1)
}

func Success(format string, a ...interface{}) {
    fmt.Print(FormatSuccess(format, a...))
}

func FormatSuccess(format string, a ...interface{}) string{
	return FormatLog(SUCCESS, format, a...)
}


func Info(format string, a ...interface{}) {
    fmt.Print(FormatInfo( format, a...))
}

func FormatInfo( format string, a ...interface{}) string{
	return FormatLog(INFO, format, a...)
}
