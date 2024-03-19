package provision

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"src/api"
	"src/utils"
)

var prefix = "provision/"

type CreateAdminDTO struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func createAdmin(c *gin.Context) {
	admin, err := utils.ToStructOrSendError[CreateAdminDTO](c)
	if err != nil {
		return
	}
	log.Debug().Msgf("Data %+v", admin)
}

func setRouters() {
	api.Router().POST(prefix+"setup", createAdmin)
}
