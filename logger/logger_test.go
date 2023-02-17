package logger

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"
)

func TestLogFormatterParamsFormat(t *testing.T) {
	type fields struct {
		TimeStamp       time.Time
		TimestampFormat string
		Level           string
		Message         string
	}
	type args struct {
		entry *logrus.Entry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LogFormatterParams{
				TimeStamp:       tt.fields.TimeStamp,
				TimestampFormat: tt.fields.TimestampFormat,
				Level:           tt.fields.Level,
				Message:         tt.fields.Message,
			}
			got, err := f.Format(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogFormatterParamsWriteCaller(t *testing.T) {
	type fields struct {
		TimeStamp       time.Time
		TimestampFormat string
		Level           string
		Message         string
	}
	type args struct {
		b     *bytes.Buffer
		entry *logrus.Entry
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LogFormatterParams{
				TimeStamp:       tt.fields.TimeStamp,
				TimestampFormat: tt.fields.TimestampFormat,
				Level:           tt.fields.Level,
				Message:         tt.fields.Message,
			}
			f.writeCaller(tt.args.b, tt.args.entry)
		})
	}
}

func TestGetColorByLevel(t *testing.T) {
	type args struct {
		level logrus.Level
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColorByLevel(tt.args.level); got != tt.want {
				t.Errorf("getColorByLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
