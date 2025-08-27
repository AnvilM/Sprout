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



func Log(level Level, format string, a ...interface{}) {
    message := fmt.Sprintf(format, a...)
    switch level {
    case INFO:
        fmt.Printf("%s[INFO]%s %s\n", Blue, Reset, message)
    case WARN:
        fmt.Printf("%s[WARN]%s %s\n", Yellow, Reset, message)
    case ERROR:
        fmt.Printf("%s[ERROR]%s %s\n", Red, Reset, message)
	case SUCCESS: 
		fmt.Printf("%s✔%s %s\n", Green, Reset, message)
    }
}

func Fatal(format string, a ...interface{}) {
    Log(ERROR, format, a...)
    os.Exit(1)
}
