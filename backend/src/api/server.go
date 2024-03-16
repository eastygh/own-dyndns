package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"src/config"
	"time"
)

type Server struct {
	gin      *gin.Engine
	defGroup *gin.RouterGroup
	server   *http.Server
}

var ginServer *Server = nil

func createEngine() *gin.Engine {
	g := gin.Default()
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	return g
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server Shutdown: %s", err)
	}
	log.Info().Msg("Server shutdown complete")
	// catching ctx.Done(). timeout of 5 seconds.
	//select {
	//case <-ctx.Done():
	//	log.Println("timeout of 5 seconds.")
	//}
	//log.Println("Server exiting")
}

// Router return default router group starting with prefix - BasePath
func Router() *gin.RouterGroup {
	return ginServer.defGroup
}

// GetServer return Server struct
func GetServer() *Server {
	return ginServer
}

func Create() {
	ginServer = &Server{}
	conf := config.Get().Server
	ginServer.gin = createEngine()
	ginServer.defGroup = ginServer.gin.Group(conf.BasePath)
	ginServer.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: ginServer.gin,
	}
}

func Start() {
	if ginServer == nil {
		log.Fatal().Msg("No API server created")
		panic("No API Server")
	}
	go func() {
		err := ginServer.server.ListenAndServe()
		if err != nil {
			return
		}
	}()
}
