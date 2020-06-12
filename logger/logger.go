package logger

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var logLevel logrus.Level
var primaryOutStream io.Writer
var LOG_PATH string = path.Join(".", "log")
var appServerLoggerFile *os.File = nil
var (
	appLogHook   logrus.Hook
	warnLogHook  logrus.Hook
	errorLogHook logrus.Hook
	coreLogHook  logrus.Hook
)

func init() {
	setLogDirectoryPath()
	setPrimaryOutStream()

	/* General hooks */
	appLogHook = lfshook.NewHook(getPathMap(path.Join(LOG_PATH, path.Base(os.Args[0])+".log")),
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true})

	warnLogHook = lfshook.NewHook(lfshook.PathMap{
		logrus.WarnLevel:  path.Join(LOG_PATH, "warn.log"),
		logrus.ErrorLevel: path.Join(LOG_PATH, "warn.log"),
		logrus.FatalLevel: path.Join(LOG_PATH, "warn.log"),
		logrus.PanicLevel: path.Join(LOG_PATH, "warn.log"),
	}, &prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true})

	errorLogHook = lfshook.NewHook(lfshook.PathMap{
		logrus.ErrorLevel: path.Join(LOG_PATH, "error.log"),
		logrus.FatalLevel: path.Join(LOG_PATH, "error.log"),
		logrus.PanicLevel: path.Join(LOG_PATH, "error.log"),
	}, &prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true})

	/* Module hooks */
	coreLogHook = lfshook.NewHook(getPathMap(path.Join(LOG_PATH, "core.log")),
		&prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true})
}

func setLogDirectoryPath() {
	appServerEnv := os.Getenv("APP_SERVER_ENV")
	if runtime.GOOS != "linux" || strings.Contains(strings.ToUpper(appServerEnv), "DEBUG") {
		LOG_PATH = path.Join(".", "log")
	} else {
		LOG_PATH = path.Join("/", "var", "log", path.Base(os.Args[0]))
	}
	if _, err := os.Stat(LOG_PATH); os.IsNotExist(err) {
		if err := os.MkdirAll(LOG_PATH, 0755); err != nil {
			logrus.Warnln(err)
		} else {
			logrus.Infoln("Directory", LOG_PATH, "is now created")
		}
	}
	switch strings.ToUpper(appServerEnv) {
	case "DEBUG":
	default:
		fullLogPath := path.Join(LOG_PATH, path.Base(os.Args[0])+".log")
		file, err := os.OpenFile(fullLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			logrus.Warnln("Error opening file.", err)
		} else {
			logrus.Infoln("Writing logs to", fullLogPath)
		}
		appServerLoggerFile = file
	}
}

func setPrimaryOutStream() {
	appServerEnv := os.Getenv("APP_SERVER_ENV")
	switch strings.ToUpper(appServerEnv) {
	case "DEBUG":
		logLevel = logrus.DebugLevel
		primaryOutStream = os.Stdout
	default:
		logLevel = logrus.InfoLevel
		primaryOutStream = ioutil.Discard // abandone output
	}
}

func getPathMap(logPath string) lfshook.PathMap {
	return lfshook.PathMap{
		logrus.DebugLevel: logPath,
		logrus.InfoLevel:  logPath,
		logrus.WarnLevel:  logPath,
		logrus.ErrorLevel: logPath,
		logrus.FatalLevel: logPath,
		logrus.PanicLevel: logPath,
	}
}

func GetLogger(module string) *logrus.Entry {
	logger := logrus.New()
	logger.Formatter = &prefixed.TextFormatter{FullTimestamp: true, ForceFormatting: true}
	logger.SetLevel(logLevel)
	setPrimaryOutStream()
	logger.AddHook(appLogHook)
	logger.AddHook(warnLogHook)
	logger.AddHook(errorLogHook)
	switch strings.ToLower(module) {
	case "core", "main":
		logger.AddHook(coreLogHook)
	}
	return logger.WithField("prefix", module)
}
