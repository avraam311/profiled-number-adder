package logger

import (
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

type Zerolog = zerolog.Logger

var Logger Zerolog

func Init(debug bool) {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Setup rolling log file with lumberjack for safe rotation
	fileWriter := &lumberjack.Logger{
		Filename:   "/app/logs/app.log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // compress old logs
	}

	multi := io.MultiWriter(fileWriter, consoleWriter)

	Logger = zerolog.New(multi).With().Timestamp().Logger()

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
}
