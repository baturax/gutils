package logger

import (
	"sync"
)

const (
	cReset  = "\033[0m"
	cGray   = "\033[90m"
	cBlue   = "\033[34m"
	cYellow = "\033[33m"
	cRed    = "\033[31m"
	cGreen  = "\033[32m"

	eDebug = ""
	eInfo  = ""
	eWarn  = ""
	eError = ""

	timeStr = "2006/01/02 15:04:05"
)

type SBaiHandler struct {
	mutex *sync.Mutex
}
