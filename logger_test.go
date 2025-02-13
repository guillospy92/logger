package logger

import (
	"log/slog"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	err := os.Setenv("LOGGER_SAVE_FILE", "true")
	if err != nil {
		t.Errorf("Error setting LOGGER_SAVE_FILE %v", err)
	}
	logger := Log()
	logger.Error("error message", slog.String("key", "value"))
}
