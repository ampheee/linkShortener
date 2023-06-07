package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var Once sync.Once
var logger zerolog.Logger

func GetLogger() zerolog.Logger {
	Once.Do(func() {
		c := color.New(color.FgRed)
		logger = zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC822,
				FormatCaller: func(i interface{}) string {
					return "|" + filepath.Base(fmt.Sprintf("%s|", i))
				},
				FormatErrFieldName: func(i interface{}) string {
					return c.Sprintf("| -> ")
				},
			}).Level(zerolog.InfoLevel).With().Timestamp().Caller().Logger()
	})
	return logger
}
