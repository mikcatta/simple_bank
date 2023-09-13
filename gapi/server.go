package gapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mikcatta/simple_bank/pb"
	"github.com/mikcatta/simple_bank/token"
	"github.com/mikcatta/simple_bank/util"
	"github.com/mikcatta/simple_bank/worker"

	db "github.com/mikcatta/simple_bank/db/sqlc"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetric)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}

/*
	func (server *Server) setupRouter() {
		router := gin.Default()

		router.POST("/users", server.createUser)
		router.POST("/users/login", server.loginUser)
		router.POST("/tokens/renew_access", server.renewAccessToken)

		authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

		authRoutes.POST("/accounts", server.createAccount)
		authRoutes.GET("/accounts/:id", server.getAccount)
		authRoutes.GET("/accounts", server.listAccounts)

		authRoutes.POST("/transfers", server.createTransfer)

		server.router = router
	}

	func (server *Server) Start(address string) error {
		log.Printf("address .... %s", address)
		return server.router.Run(address)
	}
*/
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
