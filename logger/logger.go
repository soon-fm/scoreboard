// Provides a global logger and support for creating seperate loggers

package logger

import (
	"io/ioutil"
	"strings"

	"scoreboard/logger/formatters"
	"scoreboard/logger/hooks"

	"github.com/sirupsen/logrus"
)

// Global logger any package can call without needing to construct
// a new logger for simplicity
var global *logger

// Constructs our global logger
func init() {
	global = New()
}

// A shortform wrapper around logrus.Fields
type F logrus.Fields

// Logger configuration interface
type Config interface {
	Level() string
	Format() string
	LogFile() string
	ConsoleOutput() bool
	LogstashType() string
}

// Logger
type logger struct {
	config Config
	entry  *logrus.Entry
}

// Sets up the logger according to configuration
func Setup(config Config) { global.Setup(config) }
func (l *logger) Setup(config Config) {
	l.config = config
	l.SetLevel(config.Level())
	l.ConsoleOutput(config.ConsoleOutput())
	l.LogToFile(config.LogFile())
	l.SetFormat(config.Format())
}

// Set the log level of the logger
func SetLevel(lvl string) { global.SetLevel(lvl) }
func (l *logger) SetLevel(lvl string) {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case "debug":
		l.entry.Logger.Level = logrus.DebugLevel
	case "warn":
		l.entry.Logger.Level = logrus.WarnLevel
	case "error":
		l.entry.Logger.Level = logrus.ErrorLevel
	default:
		l.entry.Logger.Level = logrus.InfoLevel
	}
}

// Enable or disbale console output
func ConsoleOutput(enable bool) {}
func (l *logger) ConsoleOutput(enable bool) {
	l.entry.Logger.Out = ioutil.Discard
	if enable {
		l.entry.Logger.Hooks.Add(hooks.NewConsoleHook())
	}
}

// Log to a file
func LogToFile(path string) { global.LogToFile(path) }
func (l *logger) LogToFile(path string) {
	if path != "" {
		l.entry.Logger.Hooks.Add(hooks.NewFileHook(path))
	}
}

// Set the format of the logger
func SetFormat(fmt string) { global.SetFormat(fmt) }
func (l *logger) SetFormat(fmt string) {
	switch fmt {
	case "logstash":
		l.setLogstashFormat(l.config.LogstashType())
	case "json":
		l.setJSONFormat()
	default:
		l.setTextFormat()
	}
}

// Set log format to use logstash format
func (l *logger) setLogstashFormat(typ string) {
	if typ == "" {
		typ = "scoreboard"
	}
	l.entry.Logger.Formatter = &formatters.LogstashFormatter{
		Type: typ,
	}
}

// Set log format to use json
func (l *logger) setJSONFormat() {
	l.entry.Logger.Formatter = &logrus.JSONFormatter{}
}

// Set text logger
func (l *logger) setTextFormat() {
	l.entry.Logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
}

// Log a field and value
func WithField(k string, v interface{}) *logger { return global.WithField(k, v) }
func (l *logger) WithField(k string, v interface{}) *logger {
	return &logger{
		config: l.config,
		entry:  l.entry.WithField(k, v),
	}
}

// Log a with multiple fields
func WithFields(fields F) *logger { return global.WithFields(fields) }
func (l *logger) WithFields(fields F) *logger {
	return &logger{
		config: l.config,
		entry:  l.entry.WithFields(logrus.Fields(fields)),
	}
}

// Log an error
func WithError(err error) *logger { return global.WithError(err) }
func (l *logger) WithError(err error) *logger {
	return &logger{
		config: l.config,
		entry:  l.entry.WithError(err),
	}
}

// Log a debug message
func Debug(msg string, v ...interface{}) { global.Debug(msg, v...) }
func (l *logger) Debug(msg string, v ...interface{}) {
	l.entry.Debugf(msg, v...)
}

// Log an info message
func Info(msg string, v ...interface{}) { global.Info(msg, v...) }
func (l *logger) Info(msg string, v ...interface{}) {
	l.entry.Infof(msg, v...)
}

// Log a warning message
func Warn(msg string, v ...interface{}) { global.Warn(msg, v...) }
func (l *logger) Warn(msg string, v ...interface{}) {
	l.entry.Warnf(msg, v...)
}

// Log an error message
func Error(msg string, v ...interface{}) { global.Error(msg, v...) }
func (l *logger) Error(msg string, v ...interface{}) {
	l.entry.Errorf(msg, v...)
}

// Log a fatal error, this causes the application to exit
func Fatal(msg string, v ...interface{}) { global.Fatal(msg, v...) }
func (l *logger) Fatal(msg string, v ...interface{}) {
	l.entry.Fatalf(msg, v...)
}

// Log a panic error, this causes the application to panic
func Panic(msg string, v ...interface{}) { global.Panic(msg, v...) }
func (l *logger) Panic(msg string, v ...interface{}) {
	l.entry.Panicf(msg, v...)
}

// Exported logger constructor, requiring a config type that
// implments the config interface
func New() *logger {
	return &logger{
		entry: logrus.NewEntry(logrus.New()),
	}
}
