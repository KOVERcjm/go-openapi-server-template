package middleware

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	/**
	 *  Colourful HTTP logger for console use
	 */
	Logger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "15:04:05",
	})
	/**
	 *  JSON format HTTP logger for log collecting use
	 */
	//Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetOutput(os.Stdout)
}
