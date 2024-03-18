package log

import (
	"fmt"
	"github.com/MR5356/go-template/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"runtime"
)

func init() {
	logrus.SetReportCaller(true)

	if config.Current().Server.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        false,
		FieldsOrder:     []string{"level"},
		TimestampFormat: "2006-01-02 15:04:05",
		TrimMessages:    true,
		CallerFirst:     true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return fmt.Sprintf(" %s:%d", frame.Function, frame.Line)
		},
	})
}
