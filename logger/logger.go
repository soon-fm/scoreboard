// Provides a global logger and support for creating seperate loggers

package logger

import (
	"io/ioutil"
	"strings"
	"time"

	"scoreboard/logger/formatters"
	"scoreboard/logger/hooks"
	"scoreboard/version"

	"github.com/sirupsen/logrus"
)

// Global logrus logger
var log = logrus.New()

// Global logger any package can call without needing to construct
// a new logger for simplicity
var global = New()

// Logger default setup
func init() {
	global.ConsoleOutput(true)
	global.SetFormat("text")
}

// A shortform wrapper around logrus.Fields
type F logrus.Fields

// Logger configuration interface
type Configurer interface {
	Level() string
	Format() string
	LogFile() string
	ConsoleOutput() bool
	LogstashType() string
}

// Logger
type logger struct {
	config Configurer
	entry  *logrus.Entry
}

// Sets up the logger according to configuration
func Setup(config Configurer) { global.Setup(config) }
func (l *logger) Setup(config Configurer) {
	l.config = config
	l.DeleteHooks()
	l.SetLevel(config.Level())
	l.ConsoleOutput(config.ConsoleOutput())
	l.LogToFile(config.LogFile())
	l.SetFormat(config.Format())
}

// Removes all hooks from the logger, call this before each setup
func (l *logger) DeleteHooks() {
	for k, _ := range log.Hooks {
		delete(log.Hooks, k)
	}
}

// Set the log level of the logger
func SetLevel(lvl string) { global.SetLevel(lvl) }
func (l *logger) SetLevel(lvl string) {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case "debug":
		log.Level = logrus.DebugLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	default:
		log.Level = logrus.InfoLevel
	}
}

// Enable or disbale console output
func ConsoleOutput(enable bool) {}
func (l *logger) ConsoleOutput(enable bool) {
	l.entry.Logger.Out = ioutil.Discard
	if enable {
		log.Hooks.Add(hooks.NewConsoleHook())
	}
}

// Log to a file
func LogToFile(path string) { global.LogToFile(path) }
func (l *logger) LogToFile(path string) {
	if path != "" {
		log.Hooks.Add(hooks.NewFileHook(path))
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
	log.Formatter = &formatters.LogstashFormatter{
		Type: typ,
	}
}

// Set log format to use json
func (l *logger) setJSONFormat() {
	log.Formatter = &logrus.JSONFormatter{}
}

// Set text logger
func (l *logger) setTextFormat() {
	log.Formatter = &logrus.TextFormatter{
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

// Log build details
func buildFields() logrus.Fields {
	v, err := version.Version()
	if err != nil {
		v = "unknown"
	}
	var tStr string
	t, err := version.BuildTime()
	if err == nil {
		tStr = t.Format(time.RFC3339)
	} else {
		tStr = "unknown"
	}
	return logrus.Fields{
		"version":   v,
		"buildTime": tStr,
	}
}

// Exported logger constructor, requiring a config type that
// implments the config interface
func New() *logger {
	return &logger{
		entry: logrus.NewEntry(log).WithFields(buildFields()),
	}
}
