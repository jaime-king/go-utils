package log

import (
	"io"
	"os"

	"github.com/jaime-king/go-utils/env"
	"github.com/sirupsen/logrus"
)

var Write *logrus.Logger

func init() {
    // Create a new logger with a JSON formatter
    Write = logrus.New()
    Write.SetFormatter(&logrus.JSONFormatter{})

    // Create the log folder
    if err := os.MkdirAll("logs", 0755); err != nil {
        Write.WithError(err).Error("Failed to create logs directory")
    }

    // Set the output to Stdout and a file
    file, err := os.OpenFile("logs/logfile.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        Write.WithError(err).Error("Failed to create log file")
    }
    
	// Load environment variables 
	env.Load()

    // Get logging level from the .env file, or default to DEBUG
    level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
    if err != nil {
        level = logrus.DebugLevel
        Write.WithError(err).Warn("Log level not correctly defined in .env file, defaulting to DEBUG")
    }
    Write.SetLevel(level)
    Write.SetOutput(io.MultiWriter(os.Stdout, file))  
}