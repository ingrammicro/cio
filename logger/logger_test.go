package logger

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"testing"
)

func TestDebugFuncInfo(t *testing.T) {
	tests := map[string]struct {
		logLevel log.Level
	}{
		"if logging level designates finer-grained informational events than the Debug": {
			logLevel: log.TraceLevel,
		},
		"if logging level lower than debug": {
			logLevel: log.InfoLevel,
		},
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			log.SetLevel(test.logLevel)
			DebugFuncInfo()

			message := buf.String()
			if !strings.Contains(message, "func logger.TestDebugFuncInfo") {
				t.Errorf("Invalid debug message: %v\n", message)
			}
		})
	}
}
