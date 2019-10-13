package logs

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// InitDefault inits logs with default values (stdout, infolevel, jsonformatter)
func InitDefault() {
	logrus.SetFormatter(&CustomJSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

const defaultTimestampFormat = time.RFC3339

type fieldKey string

// FieldMap allows customization of the key names for default fields.
type FieldMap map[fieldKey]string

// Default key names for the default fields
const (
	FieldKeyMsg   = "message"
	FieldKeyLevel = "level"
	FieldKeyTime  = "time"
)

func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}

	return string(key)
}

// CustomJSONFormatter formats logs into parsable json
type CustomJSONFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &JSONFormatter{
	//   	FieldMap: FieldMap{
	// 		 FieldKeyTime: "@timestamp",
	// 		 FieldKeyLevel: "@level",
	// 		 FieldKeyMsg: "@message",
	//    },
	// }
	FieldMap FieldMap

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

// Format renders a single log entry
func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	prefixFieldClashes(data, f.FieldMap)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	if !f.DisableTimestamp {
		data[f.FieldMap.resolve(FieldKeyTime)] = entry.Time.Format(timestampFormat)
	}
	data[f.FieldMap.resolve(FieldKeyMsg)] = entry.Message
	data[f.FieldMap.resolve(FieldKeyLevel)] = entry.Level.String()

	var serialized []byte
	var err error

	if f.PrettyPrint {
		serialized, err = json.MarshalIndent(data, "", "  ")
	} else {
		serialized, err = json.Marshal(data)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func prefixFieldClashes(data logrus.Fields, fieldMap FieldMap) {
	timeKey := fieldMap.resolve(FieldKeyTime)
	if t, ok := data[timeKey]; ok {
		data["fields."+timeKey] = t
		delete(data, timeKey)
	}

	msgKey := fieldMap.resolve(FieldKeyMsg)
	if m, ok := data[msgKey]; ok {
		data["fields."+msgKey] = m
		delete(data, msgKey)
	}

	levelKey := fieldMap.resolve(FieldKeyLevel)
	if l, ok := data[levelKey]; ok {
		data["fields."+levelKey] = l
		delete(data, levelKey)
	}
}

// Infof ...
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Info ...
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Error ...
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
