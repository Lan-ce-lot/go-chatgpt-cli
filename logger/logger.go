package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37
)

var LanLogger = logrus.New()

const (
	defaultTimestampFormat = "2006-01-02 15:04:05"
)

type LogFormatterParams struct {
	// TimeStamp is the time at which the log entry was created.
	TimeStamp time.Time
	// TimeStampFormat is the format of the timestamp.
	TimestampFormat string
	// Level is the logging level.
	Level string
	// Message is the user-supplied message.
	Message string
}

func getColorByLevel(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return colorGray
	case logrus.WarnLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorBlue
	}
}
func (f *LogFormatterParams) writeCaller(b *bytes.Buffer, entry *logrus.Entry) {
	if entry.HasCaller() {
		_, err := fmt.Fprintf(b, " [%s:%d %s]", // get filename without path
			filepath.Base(entry.Caller.File), //entry.Caller.File,
			entry.Caller.Line, entry.Caller.Function)
		if err != nil {
			return
		}
	}
}
func (f *LogFormatterParams) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	// output buffer
	b := &bytes.Buffer{}

	// write level
	level := strings.ToUpper(entry.Level.String())
	newLog := fmt.Sprintf("[%s] %s\n", level, entry.Message)

	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d] %65s\n", entry.Time.Format(timestampFormat), level, fName, entry.Caller.Line, entry.Message)
	}
	//f.writeCaller(b, entry)
	_, err := fmt.Fprintf(b, "\033[%dm%s\033[0m", levelColor, newLog)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func init() {
	//log := logrus.New()
	LanLogger.Out = os.Stdout

	LanLogger.SetLevel(logrus.DebugLevel)
	LanLogger.SetReportCaller(true)
	LanLogger.SetFormatter(&LogFormatterParams{})
}
