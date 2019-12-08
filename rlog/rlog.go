package rlog

import (
	"github.com/rs-ink/rslog"
)

const ProjectName = "wx-server"

func init() {
	SetProjectName("wx-server")
}

func CheckShowError(err error) {
	if err != nil {
		rslog.Out(1, rslog.LevelERROR, err)
	}
}

func Debug(v ...interface{}) {
	rslog.Out(1, rslog.LevelDEBUG, v...)
}

func Info(v ...interface{}) {
	rslog.Out(1, rslog.LevelINFO, v...)
}

func Warn(v ...interface{}) {
	rslog.Out(1, rslog.LevelWARN, v...)
}

func Error(v ...interface{}) {
	rslog.Out(1, rslog.LevelERROR, v...)
}

func DebugF(f string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelDEBUG, f, v...)
}

func InfoF(f string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelINFO, f, v...)
}

func WarnF(f string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelWARN, f, v...)
}

func ErrorF(f string, v ...interface{}) {
	rslog.OutF(1, rslog.LevelERROR, f, v...)
}

func SetProjectName(name string) {
	rslog.SetProjectName(name)
}

func SetRLevel(level rslog.RLevel) {
	rslog.SetRLevel(level)
}

func SetRootRLevel(level rslog.RLevel) {
	rslog.SetRootRLevel(level)
}
