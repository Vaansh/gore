package gcloud

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"github.com/Vaansh/gore/internal/config"
	"google.golang.org/api/option"
	"log"
	"os"
)

const (
	LogDirectory = "log"
)

var (
	client      *logging.Client
	cloudLogger *logging.Logger

	// local logging for development
	localWarningLogger *log.Logger
	localInfoLogger    *log.Logger
	localErrorLogger   *log.Logger
)

// Helper function to create a log file.
func openLogFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

// Initializes local loggers.
func initLocalLoggers() error {
	file, err := openLogFile(fmt.Sprintf("%s/%s.log", LogDirectory, "info"))
	if err != nil {
		return err
	}
	localInfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	file, err = openLogFile(fmt.Sprintf("%s/%s.log", LogDirectory, "warning"))
	if err != nil {
		return err
	}
	localWarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	file, err = openLogFile(fmt.Sprintf("%s/%s.log", LogDirectory, "error"))
	if err != nil {
		return err
	}
	localErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Log levels constants.
const (
	Info      = logging.Info
	Error     = logging.Error
	Warning   = logging.Warning
	Emergency = logging.Emergency
)

// InitLogger Initializes the logger based on configuration.
func InitLogger() error {
	cfg := config.ReadLoggerConfig()

	if cfg.LocalLog {
		err := initLocalLoggers()
		if err != nil {
			return err
		}
	}

	if cfg.CloudLog {
		ctx := context.Background()
		var err error
		client, err = logging.NewClient(ctx, cfg.ProjectId, option.WithCredentialsFile(cfg.CredentialsPath))
		if err != nil {
			return err
		}
		cloudLogger = client.Logger(cfg.LogName)
	}

	return nil
}

// Logs to cloud logger.
func cloudLog(severity logging.Severity, format string) {
	if cloudLogger != nil {
		logEntry := logging.Entry{
			Payload:  format,
			Severity: severity,
		}
		cloudLogger.Log(logEntry)
	}
}

// LogInfo logs info message to both local and cloud loggers.
func LogInfo(format string) {
	if localInfoLogger != nil {
		localInfoLogger.Println(format)
	}

	if cloudLogger != nil {
		cloudLog(Info, format)
	}
}

// LogError logs error message to both local and cloud loggers.
func LogError(format string) {
	if localErrorLogger != nil {
		localErrorLogger.Println(format)
	}

	if cloudLogger != nil {
		cloudLog(Error, format)
	}
}

// LogWarning logs warning message to both local and cloud loggers.
func LogWarning(format string) {
	if localWarningLogger != nil {
		localWarningLogger.Println(format)
	}

	if cloudLogger != nil {
		cloudLog(Warning, format)
	}
}

// LogFatal logs fatal message to both local and cloud loggers.
func LogFatal(format string) {
	if localErrorLogger != nil {
		localErrorLogger.Println(format)
	}

	if cloudLogger != nil {
		cloudLog(Emergency, format)
	}

	log.Fatalf(format)
}
