package logger

import (
	"log/slog"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

var DynamicLoggingConfig dynamicTimeEncoder

type dynamicTimeEncoder struct {
	encoding string
	level    zap.AtomicLevel
}

/*
SetTimeEncoder dynamically sets the time encoder at runtime:
'rfc3339': The encoder is set to a custom RFC3339 encoder
All other values: The encoder is set to an Epoch encoder
*/
func (e *dynamicTimeEncoder) SetTimeEncoder(enc string) {
	e.encoding = enc
}

func (e *dynamicTimeEncoder) encodeTime(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
	switch e.encoding {
	case "rfc3339":
		RFC3339Formatter()(t, pae)
	default:
		zapcore.EpochTimeEncoder(t, pae)
	}
}

/*
SetLoggingLevel dynamically sets the logging level at runtime. See https://github.com/uber-go/zap/blob/5786471c1d41c255c1d8b63ad30a82b68eda2c21/zapcore/level.go#L180
for possible logging levels.
*/
func (e *dynamicTimeEncoder) SetLoggingLevel(level string) {
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	e.level.SetLevel(zapLevel)
}

/*
CreateNewLogger is used to create a pre-configured slog.Logger with a zapslog handler and provided logging level,
timestamp format and writeSyncer.
*/
func CreateNewLogger(level string, timestampFormat string, writeSyncer zapcore.WriteSyncer) *slog.Logger {
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	DynamicLoggingConfig = dynamicTimeEncoder{encoding: timestampFormat, level: zap.NewAtomicLevelAt(zapLevel)}

	zapConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "log_level",
		EncodeLevel:   numberLevelFormatter,
		TimeKey:       "timestamp",
		EncodeTime:    DynamicLoggingConfig.encodeTime,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		StacktraceKey: "stack_trace",
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig),
		writeSyncer,
		DynamicLoggingConfig.level,
	)

	zapHandler := zapslog.NewHandler(zapCore, &zapslog.HandlerOptions{AddSource: true})
	slogFrontend := slog.New(zapHandler)
	return slogFrontend

}

/*
ErrAttr is creating an slog.String attribute with 'error' key and the provided error message as value.
*/
func ErrAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func numberLevelFormatter(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt(levelNumber(level))
}

// We add 1 to zap's default values to match our setLoggingLevel definitions
// https://github.com/uber-go/zap/blob/5786471c1d41c255c1d8b63ad30a82b68eda2c21/zapcore/level.go#L37
func levelNumber(level zapcore.Level) int {
	return int(level) + 1
}

/*
AppendSource returns a copy of the provided logger, which comes with the 'source' attribute set to the provided
prefix and component.
*/
func AppendSource(logger *slog.Logger, prefix string, component string) *slog.Logger {
	return logger.With(slog.String("source", prefix+"."+component))
}

/*
Panic logs message and slogAttrs with Error level. For compatibility with zlog, the function is panicking after
writing the log message.
*/
func Panic(logger *slog.Logger, message string, slogAttrs ...any) {
	logger.Error(message, slogAttrs...)
	panic(message)
}

/*
Fatal logs message and slogAttrs with Error level. For compatibility with zlog, the process is terminated
via os.Exit(1) after writing the log message.
*/
func Fatal(logger *slog.Logger, message string, slogAttrs ...any) {
	logger.Error(message, slogAttrs...)
	os.Exit(1)
}

// RFC3339Formatter TimeEncoder for RFC3339 with trailing Z for UTC and nanoseconds
func RFC3339Formatter() zapcore.TimeEncoder {
	return zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000000Z")
}
