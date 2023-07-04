package server

import (
	"context"

	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
	"github.com/modaniru/tgf-gRPC/src/service"
)

type TgfServer struct {
	pkg.TwitchGeneralFollowsServer
	service *service.Service
}

func NewServer(service *service.Service) *TgfServer{
	return &TgfServer{
		service: service,
	}
}

func (t *TgfServer) GetGeneralFollows(c context.Context, request *pkg.GetTGFRequest) (*pkg.GetTGFResponse, error){
	return t.service.GetGeneralFollows(request.GetUsernames())
}