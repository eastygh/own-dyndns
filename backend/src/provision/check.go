package provision

import (
	"github.com/rs/zerolog/log"
	"src/repository"
	"src/utils"
)

var randomInitToken string = utils.RandomString(30)

func setRouters() {

}

func DoProvision() {
	i, err := repository.GetIntParameter("init")
	if err != nil {
		log.Error().Err(err)
		panic("Can't instance check")
	}
	if i == 0 {
		// Do init job
		setRouters()
		log.Info().Msgf("Initialization token %s", randomInitToken)
	} else {
		log.Info().Msg("The service has been initialized")
	}
}
