package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Server struct {
	gin    *gin.Engine
	server *http.Server
}

var GinServer *Server

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

func Create() {
	GinServer = &Server{}
	GinServer.gin = gin.Default()
	GinServer.server = &http.Server{
		Addr:    ":8080",
		Handler: GinServer.gin,
	}
}

func Start() {
	if GinServer == nil {
		log.Fatal().Msg("No API server created")
		panic("No API Server")
	}
	go func() {
		err := GinServer.server.ListenAndServe()
		if err != nil {
			return
		}
	}()
}
