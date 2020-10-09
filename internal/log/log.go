package log

import (
	"base-site-api/internal/app/config"
	"fmt"
	"io"
	"os"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

// init setup logger
func Setup(c *config.ApplicationConfiguration) {
	var logLevel log.Level

	logger = log.New()

	logger.SetFormatter(&log.TextFormatter{})

	if c.Debug {
		logLevel = log.DebugLevel
	} else {
		logLevel = log.InfoLevel
	}

	logger.SetLevel(logLevel)

	if c.LogToFile {
		file, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("Fatal while opening out: %s", err)
		}

		logger.SetOutput(io.MultiWriter(os.Stdout, file))
		log.RegisterExitHandler(func() {
			if file == nil {
				return
			}
			err := file.Close()
			if err != nil {
				log.Debugf("Problem with closing file %s", err)
			}
		})
	} else {
		logger.SetOutput(os.Stdout)
	}

	initSentry(c.SentryDNS, c.Env)
}

func initSentry(dns string, env string) {
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dns,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug:       false,
		Environment: env,
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
