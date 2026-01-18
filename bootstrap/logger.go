package bootstrap

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// ANSI color codes for log levels
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// LogFormatter implements logrus.Formatter with colored output
type LogFormatter struct{}

// Format renders a single log entry
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// Format: [timestamp] [LEVEL] file:line func message
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	// Include caller info if available
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s]\x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s]\x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}

	// Return the formatted byte slice
	return b.Bytes(), nil
}

// InitLogger initializes and returns a configured logrus.Logger instance.
// It receives a Config pointer (call LoadConfig first to obtain it).
func InitLogger(conf *Config) *logrus.Logger {

	cfg := conf.Logger

	var mlog *logrus.Logger = logrus.New()
	mlog.SetOutput(os.Stdout)
	mlog.SetReportCaller(cfg.ShowLine)
	mlog.SetFormatter(&LogFormatter{})

	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mlog.SetLevel(level)

	// Also configure the global logrus logger for convenience
	initDefaultLogger(cfg)

	return mlog
}

// initDefaultLogger configures the global logrus standard logger
func initDefaultLogger(cfg Logger) {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(cfg.ShowLine)
	logrus.SetFormatter(&LogFormatter{})

	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
