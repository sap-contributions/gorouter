package logger_test

import (
	"code.cloudfoundry.org/lager/v3"
	"github.com/onsi/gomega/gbytes"
	"go.uber.org/zap/zapcore"

	goRouterLogger "code.cloudfoundry.org/gorouter/logger"
	"code.cloudfoundry.org/gorouter/test_util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LagerAdapter", func() {
	var (
		lagerLogger  lager.Logger
		testSink     *test_util.TestZapSink
		sourcePrefix string
	)

	BeforeEach(func() {
		sourcePrefix = "gorouter"
		testSink = &test_util.TestZapSink{Buffer: gbytes.NewBuffer()}
		slogLogger := goRouterLogger.CreateNewLogger("DEBUG", "", zapcore.NewMultiWriteSyncer(testSink, zapcore.AddSync(GinkgoWriter)))
		lagerLogger = goRouterLogger.NewLagerAdapter(slogLogger, sourcePrefix)
	})

	Describe("Session", func() {
		Context("when multiple sessions are appended", func() {
			var sessionString1 = "component"
			var sessionString2 = "subcomponent"
			var message = "some-action"
			It("adds the concatenated sessions as source", func() {
				lagerLogger = lagerLogger.Session(sessionString1)
				lagerLogger = lagerLogger.Session(sessionString2)
				lagerLogger.Info(message)

				Expect(testSink.Lines()).To(HaveLen(1))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":[0-9]+[.][0-9]+,"message":"%s","source":"%s[.]%s[.]%s"}`,
					message,
					sourcePrefix,
					sessionString1,
					sessionString2,
				))

			})
		})

		Context("when a session with data is appended", func() {
			var sessionString1 = "component"
			var message = "some-action"
			It("adds the session and data", func() {
				lagerLogger = lagerLogger.Session(sessionString1, lager.Data{"foo": "bar", "bar": "baz"})
				lagerLogger.Info(message)

				Expect(testSink.Lines()).To(HaveLen(1))
				Expect(testSink.Lines()[0]).To(MatchRegexp(
					`{"log_level":[0-9]*,"timestamp":[0-9]+[.][0-9]+,"message":"%s","source":"%s[.]%s","foo":"bar","bar":"baz"}`,
					message,
					sourcePrefix,
					sessionString1,
				))

			})
		})

	})
})
