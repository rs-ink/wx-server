package rlog

import (
	"xorm.io/core"
)

type XormLog struct {
	isDebug  bool
	logLevel core.LogLevel
	showSQL  bool
}

func (xl XormLog) Debug(v ...interface{}) {

}

func (xl XormLog) Debugf(format string, v ...interface{}) {

}

func (xl XormLog) Error(v ...interface{}) {

}

func (xl XormLog) Errorf(format string, v ...interface{}) {

}

func (xl XormLog) Info(v ...interface{}) {

}

func (xl XormLog) Infof(format string, v ...interface{}) {

}

func (xl XormLog) Warn(v ...interface{}) {

}

func (xl XormLog) Warnf(format string, v ...interface{}) {

}

func (xl XormLog) Level() core.LogLevel {
	return xl.logLevel
}

func (xl *XormLog) SetLevel(l core.LogLevel) {
	xl.logLevel = l
}

func (xl *XormLog) ShowSQL(show ...bool) {
	if len(show) > 0 {
		xl.showSQL = show[0]
	}
}

func (xl XormLog) IsShowSQL() bool {
	return xl.showSQL
}
