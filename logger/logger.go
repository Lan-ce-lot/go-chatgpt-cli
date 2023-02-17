package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37
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
		_, err := fmt.Fprintf(
			b,
			" [%s:%d %s]",
			// get filename without path
			filepath.Base(entry.Caller.File),
			//entry.Caller.File,
			entry.Caller.Line,
			entry.Caller.Function,
		)
		if err != nil {
			return
		}
	}
}
func (f *LogFormatterParams) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.StampMilli
	}

	// output buffer
	b := &bytes.Buffer{}

	// write level
	level := strings.ToUpper(entry.Level.String())

	f.writeCaller(b, entry)
	return []byte(fmt.Sprintf(
		"[%s] \x1b[%dm%-7s\x1b[0m %-25s%-10s\n",
		entry.Time.Format("2006/01/02 - 15:04:05"),
		levelColor,
		level,
		entry.Message,
		b.String(),
	)), nil
}

func init() {
	log := logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&LogFormatterParams{})
	log.SetLevel(logrus.DebugLevel)
}
