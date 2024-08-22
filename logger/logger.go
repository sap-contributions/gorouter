package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

// Logger is the zap.Logger interface with additional Session methods.
//
//type Logger interface {
//	With(...slog.Attr) Logger
//	Debug(string, ...slog.Attr)
//	Info(string, ...slog.Attr)
//	Warn(string, ...slog.Attr)
//	Error(string, ...slog.Attr)
//	Panic(string, ...slog.Attr)
//	Fatal(string, ...slog.Attr)
//	Session(string) Logger
//	SessionName() string
//	ErrAttr(error) slog.Attr
//}
//
//type logger struct {
//	source string
//	slog.Logger
//}

func CreateNewLogger(component string, level string, timestampFormat string) *slog.Logger {
	return CreateNewLoggerWithWriteSyncer(component, level, timestampFormat, zapcore.Lock(os.Stdout))
}

func CreateNewLoggerWithWriteSyncer(component string, level string, timestampFormat string, writeSyncer zapcore.WriteSyncer) *slog.Logger {
	var logLevel slog.Level
	logLevel.UnmarshalText([]byte(level))

	formatter := zapcore.RFC3339NanoTimeEncoder

	if timestampFormat == "rfc3339" {
		formatter = zapcore.RFC3339TimeEncoder
	}
	//opts := slog.HandlerOptions{
	//	Level: logLevel,
	//}
	//
	//lggr := slog.New(zapslog.NewHandler(zap.L().Core(), &zapslog.HandlerOptions{
	//	LoggerName: "",
	//	AddSource:  false,
	//}))

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

	//core := zapcore.NewCore(
	//	zapcore.NewJSONEncoder(zapConfig),
	//	zapcore.Lock(os.Stdout),
	//	zapLevel,
	//)

	//zapInstance := zap.New(
	//	zapcore.NewLazyWith(
	//		core,
	//		[]zapcore.Field{{
	//			Key:    "source",
	//			String: component,
	//			Type:   zapcore.StringType,
	//		}},
	//	),
	//	zap.AddStacktrace(zapcore.ErrorLevel),
	//	zap.AddCaller(),
	//)

	//zapInstance2 := zap.New(
	//	core,
	//	zap.AddStacktrace(zapcore.ErrorLevel),
	//	zap.AddCaller(),
	//).With(zapcore.Field{
	//	Key:    "source",
	//	String: component,
	//	Type:   zapcore.StringType,
	//})

	//			zap.New(zapcore.NewLazyWith(core, []zapcore.Field{{
	//	Key:    "source",
	//	String: component,
	//	Type:   zapcore.StringType,
	//}})),

	//zapInstance := zap.New(core,
	//	zap.AddCaller(),
	//	zap.AddStacktrace(zapcore.ErrorLevel),
	//)
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig),
		writeSyncer,
		zapLevel,
	)
	zapHandler := zapslog.NewHandler(zapCore, &zapslog.HandlerOptions{AddSource: true})

	slogFrontend := slog.New(zapHandler).With(slog.String("source", component))
	slog.SetDefault(slogFrontend)
	slog.SetLogLoggerLevel(logLevel)

	return slogFrontend

}

//// NewLogger returns a new zap logger that implements the Logger interface.
//func NewLogger(component string, timestampFormat string, logLevel zapcore.Level, writeSyncer zapcore.WriteSyncer) Logger {
//	formatter := zapcore.RFC3339NanoTimeEncoder
//
//	if timestampFormat == "rfc3339" {
//		formatter = zapcore.RFC3339TimeEncoder
//	}
//
//	conf := zapcore.EncoderConfig{
//		MessageKey:  "message",
//		LevelKey:    "log_level",
//		TimeKey:     "ts",
//		EncodeLevel: numberLevelFormatter,
//		EncodeTime:  formatter,
//	}
//
//	core := zapcore.NewCore(zapcore.NewJSONEncoder(conf), writeSyncer, logLevel)
//
//	origLogger := zap.New(core)
//
//	return &logger{
//		source:     component,
//		origLogger: *origLogger,
//		Logger: *zap.New(origLogger.Core()).With(zapcore.Field{
//			Key:    "source",
//			String: component,
//			Type:   zapcore.StringType,
//		}),
//	}
//}

// func (l *logger) Session(component string) Logger {
//
//		newSource := l.source + "." + component
//		return l.With(slog.String("source", newSource))
//
//		//lggr := &logger{
//		//	source: newSource,
//		//	zapCore: l.zapCore.With([]zapcore.Field{{
//		//		Key:    "source",
//		//		String: component,
//		//		Type:   zapcore.StringType,
//		//	}}),
//		//	Logger: *slog.New(zapslog.NewHandler(l.zapCore, &zapslog.HandlerOptions{AddSource: true})),
//		//	//Logger: *zap.New(l.origLogger.Core()).With(zapcore.Field{
//		//	//	Key:    "source",
//		//	//	String: newSource,
//		//	//	Type:   zapcore.StringType,
//		//	//}),
//		//	context: l.context,
//		//}
//		//return lggr
//	}
//
//	func (l *logger) SessionName() string {
//		return l.source
//	}
//
//	func (l *logger) With(fields ...slog.Attr) Logger {
//		l.Logger.With(fields)
//		//return &logger{
//		//	source:  l.source,
//		//	zapCore: l.zapCore,
//		//	Logger:  l.Logger,
//		//	context: append(l.context, fields),
//		//}
//	}
//
//	func (l *logger) Debug(msg string, fields ...slog.Attr) {
//		l.Logger.Debug(msg, fields)
//	}
//
//	func (l *logger) Info(msg string, fields ...slog.Attr) {
//		l.Logger.Info(msg, fields)
//	}
//
//	func (l *logger) Warn(msg string, fields ...slog.Attr) {
//		l.Logger.Warn(msg, fields)
//	}
//
//	func (l *logger) Error(msg string, fields ...slog.Attr) {
//		l.Logger.Error(msg, fields)
//	}
//
//	func (l *logger) Panic(msg string, fields ...slog.Attr) {
//		l.Logger.Error(msg, fields)
//		panic(msg)
//	}
//
//	func (l *logger) Fatal(msg string, fields ...slog.Attr) {
//		l.Logger.Error(msg, fields)
//		os.Exit(1)
//	}
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
	return logger.With(slog.String("source", prefix+".nats"))
}
