package middleware

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "15:04:05",
	})
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetOutput(os.Stdout)
}
