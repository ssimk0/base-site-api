package log

import (
	"fmt"
	"io"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

// init setup logger
func init() {
	var logLevel log.Level
	var err error

	logger = log.New()

	file, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Fatal while opening out: %s", err)
	}

	logger.SetFormatter(&log.TextFormatter{})

	if os.Getenv("GO_ENV") == "testing" {
		err = godotenv.Load(".test.env")
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		logger.Warnf("Fatal while loading env: %s", err)
	}

	if os.Getenv("GO_ENV") == "development" {
		logLevel = log.DebugLevel
	} else {
		logLevel = log.InfoLevel
	}

	logger.SetLevel(logLevel)
	logger.SetOutput(io.MultiWriter(os.Stdout, file))

	log.RegisterExitHandler(func() {
		if file == nil {
			return
		}
		file.Close()
	})

	initSentry()
}

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: os.Getenv("SENTRY_DNS"),
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: false,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

// Info log
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof log with format
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Debug log
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf log with format
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Error log
func Error(err error) {
	sentry.CaptureException(err)
	logger.Error(err)
}

// Errorf log with format
func Errorf(format string, args ...interface{}) {
	sentry.CaptureException(fmt.Errorf(format, args...))
	logger.Errorf(format, args...)
}

// Fatal log
func Fatal(err error) {
	sentry.CaptureException(err)
	logger.Fatal(err)
}

// Fatalf log with format
func Fatalf(format string, args ...interface{}) {
	sentry.CaptureException(fmt.Errorf(format, args...))
	logger.Fatalf(format, args...)
}

// Print log
func Print(args ...interface{}) {
	logger.Print(args...)
}

// Printf log with format
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Writer return a instance of write
func Writer() *io.PipeWriter {
	return logger.Writer()
}
