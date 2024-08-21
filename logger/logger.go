package logger

import (
	"strconv"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// Logger is the zap.Logger interface with additional Session methods.
//
//go:generate counterfeiter -o fakes/fake_logger.go . Logger
type Logger interface {
	With(...zap.Field) Logger
	Debug(string, ...zap.Field)
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, ...zap.Field)
	DPanic(string, ...zap.Field)
	Panic(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	Session(string) Logger
	SessionName() string
}

type logger struct {
	source     string
	origLogger zap.Logger
	context    []zap.Field
	zap.Logger
}

// NewLogger returns a new zap logger that implements the Logger interface.
func NewLogger(component string, timestampFormat string, logLevel zapcore.Level, writeSyncer zapcore.WriteSyncer) Logger {
	formatter := zapcore.RFC3339NanoTimeEncoder

	if timestampFormat == "rfc3339" {
		formatter = zapcore.RFC3339TimeEncoder
	}

	conf := zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "log_level",
		TimeKey:     "ts",
		EncodeLevel: numberLevelFormatter,
		EncodeTime:  formatter,
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(conf), writeSyncer, logLevel)

	origLogger := zap.New(core)

	return &logger{
		source:     component,
		origLogger: *origLogger,
		Logger: *zap.New(origLogger.Core()).With(zapcore.Field{
			Key:    "source",
			String: component,
			Type:   zapcore.StringType,
		}),
	}
}

func (l *logger) Session(component string) Logger {
	newSource := l.source + "." + component
	lggr := &logger{
		source:     newSource,
		origLogger: l.origLogger,
		Logger: *zap.New(l.origLogger.Core()).With(zapcore.Field{
			Key:    "source",
			String: newSource,
			Type:   zapcore.StringType,
		}),
		context: l.context,
	}
	return lggr
}

func (l *logger) SessionName() string {
	return l.source
}

func (l *logger) With(fields ...zap.Field) Logger {
	return &logger{
		source:     l.source,
		origLogger: l.origLogger,
		Logger:     l.Logger,
		context:    append(l.context, fields...),
	}
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}
func (l *logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}
func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.Logger.Warn(msg, fields...)
}
func (l *logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}
func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.Logger.Panic(msg, fields...)
}
func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
}

func numberLevelFormatter(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strconv.Itoa(levelNumber(level)))
}

// We add 1 to zap's default values to match our level definitions
// https://github.com/uber-go/zap/blob/47f41350ff078ea1415b63c117bf1475b7bbe72c/level.go#L36
func levelNumber(level zapcore.Level) int {
	return int(level) + 1
}
