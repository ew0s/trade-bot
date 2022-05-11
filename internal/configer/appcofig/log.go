package appcofig

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
	Fatal LogLevel = "fatal"
	Panic LogLevel = "panic"
	Trace LogLevel = "trace"
)

var logLevelNames = map[LogLevel]int{
	Trace: 6,
	Debug: 5,
	Info:  4,
	Warn:  3,
	Error: 2,
	Fatal: 1,
	Panic: 0,
}

func (l LogLevel) Int() int {
	return logLevelNames[l]
}

type LogFormat string

const (
	Text LogFormat = "text"
	JSON LogFormat = "json"
)

type Logger struct {
	Project           string    `yaml:"project"`
	Format            LogFormat `yaml:"format"`
	Level             LogLevel  `yaml:"level"`
	Env               string    `yaml:"env"`
	DisableStackTrace bool      `yaml:"disable_stack_trace"`
}
