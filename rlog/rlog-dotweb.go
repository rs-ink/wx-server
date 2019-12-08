package rlog

import "github.com/rs-ink/rslog"

func NewRAppLog() (logger *RAppLog) {
	logger = &RAppLog{
		log: rslog.NewRsLog(true),
	}
	logger.log.SetProjectName(ProjectName)
	return
}

type RAppLog struct {
	log            *rslog.RsLog
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
	rl.log.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Raw(log string, logTarget string) {
	rl.log.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Debug(log string, logTarget string) {
	rl.log.Out(1, rslog.LevelDEBUG, log)
}

func (rl RAppLog) Info(log string, logTarget string) {
	rl.log.Out(1, rslog.LevelINFO, log)
}

func (rl RAppLog) Warn(log string, logTarget string) {
	rl.log.Out(1, rslog.LevelWARN, log)
}

func (rl RAppLog) Error(log string, logTarget string) {
	rl.log.Out(1, rslog.LevelERROR, log)
}
