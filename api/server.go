package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/saram12saram2/bank/db/sqlc"
)


type Server struct {
	store  *db.Store
    router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
    return server.router.Run(address)
}
