package rlog

import (
	"github.com/rs-ink/rslog"
	"xorm.io/core"
)

type XormLog struct {
	isDebug  bool
	logLevel core.LogLevel
	showSQL  bool
}

func (xl XormLog) Debug(v ...interface{}) {
	rslog.Out(1, rslog.LevelDEBUG, v)
}

func (xl XormLog) Debugf(format string, v ...interface{}) {
	rslog.OutF(1,rslog.LevelDEBUG,format,v...)
}

func (xl XormLog) Error(v ...interface{}) {
	rslog.Out(1, rslog.LevelERROR, v)
}

func (xl XormLog) Errorf(format string, v ...interface{}) {
	rslog.OutF(1,rslog.LevelERROR,format,v...)
}

func (xl XormLog) Info(v ...interface{}) {
	rslog.Out(1, rslog.LevelINFO, v)
}

func (xl XormLog) Infof(format string, v ...interface{}) {
	rslog.OutF(1,rslog.LevelINFO,format,v...)
}

func (xl XormLog) Warn(v ...interface{}) {
	rslog.Out(1, rslog.LevelWARN, v)
}

func (xl XormLog) Warnf(format string, v ...interface{}) {
	rslog.OutF(1,rslog.LevelWARN,format,v...)
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
