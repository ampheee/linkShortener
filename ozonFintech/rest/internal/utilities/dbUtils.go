package utilities

import (
	"fmt"
	"ozonFintech/pkg/logger"
	"time"
)

func ConnectWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	log := logger.GetLogger()
	for i := 0; i < attempts; i++ {
		err = fn()
		if err != nil {
			log.Warn().Err(err).Msg(fmt.Sprintf("Try num: %d, connection failed", i))
			time.Sleep(delay)
			continue
		}
	}
	return err
}
