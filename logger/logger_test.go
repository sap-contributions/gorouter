package logger_test

import (
	"errors"
	"log/slog"

	"go.uber.org/zap/zapcore"

	log "code.cloudfoundry.org/gorouter/logger"
	"code.cloudfoundry.org/gorouter/test_util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Logger", func() {
	var logger *slog.Logger
	var testSink *test_util.TestZapSink
	var action = "my-action"

	BeforeEach(func() {
		testSink = &test_util.TestZapSink{Buffer: gbytes.NewBuffer()}
		logger = log.CreateNewLogger(
			"INFO",
			"",
			zapcore.NewMultiWriteSyncer(testSink, zapcore.AddSync(GinkgoWriter)))
	})

	Describe("CreateNewLogger", func() {
		Context("when timestampFormat is omitted", func() {
			It("outputs a properly-formatted message with epoch timestamp", func() {
				logger.Info(action)
				Expect(testSink.Lines()).To(HaveLen(1))

				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":[0-9]+[.][0-9]+,"message":"%s"}`,
					action,
				))
			})
		})

		Context("when timestampFormat is omitted", func() {
			BeforeEach(func() {
				logger = log.CreateNewLogger(
					"info",
					"rfc3339",
					zapcore.NewMultiWriteSyncer(testSink, zapcore.AddSync(GinkgoWriter)))
			})

			It("outputs a properly-formatted message with rfc3339 timestamp", func() {
				logger.Info(action)
				Expect(testSink.Lines()).To(HaveLen(1))

				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{9}Z","message":"%s"}`,
					action,
				))
			})
		})

		Context("when logs at different levels are written", func() {
			It("outputs log messages with correct numeric 'log_level' value according to 'numberLevelFormatter'", func() {
				log.DynamicLoggingConfig.SetLoggingLevel("Debug")
				logger.Debug("debug")
				logger.Info("info")
				logger.Warn("warn")
				logger.Error("error")

				Expect(testSink.Lines()).To(HaveLen(4))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":0,"timestamp":[0-9]+[.][0-9]+,"message":"debug"}`,
				))
				Expect(testSink.Lines()[1]).To(MatchRegexp(
					`{"log_level":1,"timestamp":[0-9]+[.][0-9]+,"message":"info"}`,
				))
				Expect(testSink.Lines()[2]).To(MatchRegexp(
					`{"log_level":2,"timestamp":[0-9]+[.][0-9]+,"message":"warn"}`,
				))
				Expect(testSink.Lines()[3]).To(MatchRegexp(
					`{"log_level":3,"timestamp":[0-9]+[.][0-9]+,"message":"error"}`,
				))
			})
		})

	})

	Describe("SetTimeEncoder", func() {
		Context("when timestampFormat is set to rfc3339", func() {
			It("outputs a properly-formatted message with RFC3339 timestamp", func() {
				log.DynamicLoggingConfig.SetTimeEncoder("rfc3339")
				logger.Info(action)
				Expect(testSink.Lines()).To(HaveLen(1))

				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{9}Z","message":"%s"}`,
					action,
				))
			})
		})
	})

	Describe("SetLoggingLevel", func() {
		Context("when logLevel is set to Info", func() {
			It("only outputs log messages with level Info and above", func() {
				log.DynamicLoggingConfig.SetLoggingLevel("Info")
				logger.Info("this-info-is-logged")
				logger.Debug("this-debug-is-not-logged")

				Expect(testSink.Lines()).To(HaveLen(1))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":[0-9]+[.][0-9]+,"message":"this-info-is-logged"}`,
				))
			})
		})

		Context("when logLevel is set to Debug", func() {
			It("outputs log messages with level Debug and above", func() {
				log.DynamicLoggingConfig.SetLoggingLevel("Debug")
				logger.Info("this-info-is-logged")
				logger.Debug("this-debug-is-logged")

				Expect(testSink.Lines()).To(HaveLen(2))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":[0-9]+[.][0-9]+,"message":"this-info-is-logged"}`,
				))
				Expect(testSink.Lines()[1]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":[0-9]+[.][0-9]+,"message":"this-debug-is-logged"}`,
				))
			})
		})
	})

	Describe("ErrAttr", func() {
		Context("when appending an error created by ErrAttr ", func() {
			It("outputs log messages with 'error' attribute", func() {
				err := errors.New("this-is-an-error")
				logger.Error(action, log.ErrAttr(err))

				Expect(testSink.Lines()).To(HaveLen(1))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":3,"timestamp":[0-9]+[.][0-9]+,"message":"%s","error":"%s"}`, action, err.Error(),
				))
			})
		})
	})

	Describe("AppendSource", func() {
		Context("when a source attribute is added via AppendSource", func() {
			BeforeEach(func() {
				logger = log.AppendSource(logger, "my-component", "my-subcomponent")
				logger.Info(action)
			})

			It("outputs log messages with source attribute", func() {
				Expect(testSink.Lines()).To(HaveLen(1))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":[0-9]+[.][0-9]+,"message":"%s","source":"my-component.my-subcomponent"}`,
					action,
				))
			})
		})

	})

	Describe("Panic", func() {
		Context("when an error is logged with 'Panic'", func() {

			It("outputs an error log message and panics", func() {
				Expect(func() { log.Panic(logger, action) }).To(Panic())

				Expect(testSink.Lines()).To(HaveLen(1))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":3,"timestamp":[0-9]+[.][0-9]+,"message":"%s"`,
					action,
				))
			})
		})
	})

})
