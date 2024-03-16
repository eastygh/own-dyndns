package provision

import (
	"github.com/rs/zerolog/log"
	"src/repository"
	"src/utils"
)

var randomInitToken string = utils.RandomString(30)

func isNeedToSetup(panicOnError bool) bool {
	i, err := repository.GetIntParameter("init")
	if err != nil {
		log.Error().Err(err)
		if panicOnError {
			panic("Can't instance check")
		}
		i = 0
	}
	return i == 0
}

func DoProvision() {
	if isNeedToSetup(true) {
		// Do init job
		setRouters()
		log.Info().Msgf("Initialization token %s", randomInitToken)
	} else {
		log.Info().Msg("The service has been initialized")
	}
}
