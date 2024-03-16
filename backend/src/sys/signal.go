package sys

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"src/api"
	"syscall"
)

func UntilEndOfDays() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		log.Info().Msgf("Catched signal - %s", sig)
		done <- true
	}()
	<-done
	if api.GetServer() != nil {
		api.GetServer().Stop()
	}
}
