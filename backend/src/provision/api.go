package provision

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"src/api"
)

var prefix = "provision/"

func createAdmin(c *gin.Context) {
	log.Debug().Msg("path")
}

func setRouters() {
	api.Router().POST(prefix+"setup", createAdmin)
}
