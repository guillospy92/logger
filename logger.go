package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime/debug"
	"strconv"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *slog.Logger
var once sync.Once

// Log logger function principal
func Log() *slog.Logger {
	once.Do(func() {

		loggerSaveFile := os.Getenv("LOGGER_SAVE_FILE")

		output := verifyCreateFile(loggerSaveFile)

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		jsonHandler := slog.NewJSONHandler(output, &slog.HandlerOptions{
			ReplaceAttr: replaceAttr,
			AddSource:   true,
			Level:       slog.LevelDebug,
		}).WithAttrs([]slog.Attr{
			slog.String("service", "worker"),
			slog.String("git_revision", gitRevision),
			slog.String("go_Version", buildInfo.GoVersion),
		})

		handler := &ContextHandler{jsonHandler}
		logger = slog.New(handler)
	})

	return logger
}

func verifyCreateFile(loggerSaveFile string) io.Writer {
	isSaveLoggerFile := false

	if loggerSaveFile != "" {
		var err error
		isSaveLoggerFile, err = strconv.ParseBool(loggerSaveFile)
		if err != nil {
			panic(fmt.Errorf("can not parse logger save file, please set var LOGGER_SAVE_FILE should be boolean %w", err))
		}

	}

	output := io.Writer(os.Stderr)

	if isSaveLoggerFile {
		fileLogger := &lumberjack.Logger{
			Filename:   "logs/logs.log",
			MaxSize:    20,
			MaxBackups: 10,
			MaxAge:     14,
			// Compress:   true,
		}

		output = io.MultiWriter(os.Stderr, fileLogger)
	}

	return output
}
