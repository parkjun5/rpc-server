package network

import (
	"rpc-server/config"
	"rpc-server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg *config.Config

	service *service.Service

	engin *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	n := &Network{cfg: cfg, service: service, engin: gin.New()}

	return n, nil
}

func (n *Network) StartServer() {
	n.engin.Run(":8080")
}
