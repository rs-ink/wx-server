package rlog

import "github.com/rs-ink/rslog"

func NewRAppLog() *RAppLog {
	return &RAppLog{}
}

type RAppLog struct {
	logPath        string
	enabledConsole bool
	enabledLog     bool
}

func (rl *RAppLog) SetLogPath(logPath string) {
	rl.logPath = logPath
}

func (rl *RAppLog) SetEnabledConsole(enabled bool) {
	rl.enabledConsole = enabled
}

func (rl *RAppLog) SetEnabledLog(enabledLog bool) {
	rl.enabledLog = enabledLog
}

func (rl RAppLog) IsEnabledLog() bool {
	return rl.enabledLog
}

func (rl RAppLog) Print(log string, logTarget string) {
	rslog.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Raw(log string, logTarget string) {
	rslog.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Debug(log string, logTarget string) {
	rslog.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Info(log string, logTarget string) {
	rslog.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Warn(log string, logTarget string) {
	rslog.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Error(log string, logTarget string) {
	rslog.Out(1, rslog.LevelINFO, log)
}
