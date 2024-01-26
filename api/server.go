package api

import (
	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/util"
)

type Server struct {
	config util.Config
	router *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/generate/wallet", server.createWallet)
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
