package server

import (
	"context"

	"github.com/modaniru/tgf-gRPC/internal/service"
	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
)

type TgfServer struct {
	pkg.TwitchGeneralFollowsServer
	service *service.Service
}

// Return server.TgfServer
func NewServer(service *service.Service) *TgfServer {
	return &TgfServer{
		service: service,
	}
}

// Return general follow list by request
func (t *TgfServer) GetGeneralFollows(c context.Context, request *pkg.GetTGFRequest) (*pkg.GetTGFResponse, error) {
	return t.service.GetGeneralFollows(request.GetUsernames())
}
