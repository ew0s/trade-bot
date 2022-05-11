package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/ew0s/trade-bot/internal/configer/appcofig"
)

func Setup(cfg appcofig.Logger) *logrus.Entry {
	var formatter logrus.Formatter

	switch cfg.Format {
	case appcofig.Text:
		formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "02-01-2006 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
			},
		}

	case appcofig.JSON:
		formatter = &logrus.JSONFormatter{
			TimestampFormat: "02-01-2006 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
			},
		}
	}

	logger := logrus.New()

	logger.SetReportCaller(true)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(formatter)

	return logger.WithFields(map[string]interface{}{
		"project": cfg.Project,
		"env":     cfg.Env,
	})
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
