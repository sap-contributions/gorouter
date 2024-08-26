package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"sync"

	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

var lock = sync.Mutex{}
var writeSyncer zapcore.WriteSyncer

func CreateNewLogger(level string, timestampFormat string) *slog.Logger {

	if writeSyncer == nil {
		defer lock.Unlock()
		lock.Lock()
		if writeSyncer == nil {
			writeSyncer = zapcore.Lock(os.Stdout)
		}
	}

	return CreateNewLoggerWithWriteSyncer(level, timestampFormat, zapcore.Lock(os.Stdout))
}

func CreateNewLoggerWithWriteSyncer(level string, timestampFormat string, writeSyncer zapcore.WriteSyncer) *slog.Logger {
	var logLevel slog.Level
	logLevel.UnmarshalText([]byte(level))

	formatter := zapcore.RFC3339NanoTimeEncoder

	if timestampFormat == "rfc3339" {
		formatter = zapcore.RFC3339TimeEncoder
	}

	zapConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "log_level",
		EncodeLevel:   numberLevelFormatter,
		TimeKey:       "ts",
		EncodeTime:    formatter,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		StacktraceKey: "stack_trace",
	}

	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		panic(fmt.Errorf("unknown log level: %s", level))
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig),
		writeSyncer,
		zapLevel,
	)
	zapHandler := zapslog.NewHandler(zapCore, &zapslog.HandlerOptions{AddSource: true})

	slogFrontend := slog.New(zapHandler)
	slog.SetDefault(slogFrontend)
	slog.SetLogLoggerLevel(logLevel)
	return slogFrontend

}

func ErrAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func numberLevelFormatter(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strconv.Itoa(levelNumber(level)))
}

// We add 1 to zap's default values to match our level definitions
// https://github.com/uber-go/zap/blob/47f41350ff078ea1415b63c117bf1475b7bbe72c/level.go#L36
func levelNumber(level zapcore.Level) int {
	return int(level) + 1
}

func AppendSource(logger *slog.Logger, prefix string, source string) *slog.Logger {
	return logger.With(slog.String("source", prefix+"."+source))
}

func Panic(logger *slog.Logger, message string, slogAttrs ...any) {
	logger.Error(message, slogAttrs...)
	panic(message)
}

func Fatal(logger *slog.Logger, message string, slogAttrs ...any) {
	logger.Error(message, slogAttrs...)
	os.Exit(1)
}
