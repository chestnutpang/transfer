package log

import (
	"cts/basic/logbasic"
	"cts/global"
	"cts/ibasic/logibasic"
)

const (
	TimeFormat         = "2020-01-01 11:11:11"
	FileTimeFormat     = "2020-01-01"
	RollingTimePattern = "0 0  * * *"
)

var l logibasic.ILogger

func Init() {
	l = logbasic.NewZapLogger()
	//setOutput(l, global.CONF.LogConfig)
	global.LOG = l
}
