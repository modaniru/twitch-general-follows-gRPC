package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 8080

type MyServer struct {
	pkg.TwitchGeneralFollowsServer
}

func (s *MyServer) GetGeneralFollows(c context.Context, request *pkg.GetTGFRequest) (*pkg.GetTGFResponse, error){
	fmt.Println(*request)
	return &pkg.GetTGFResponse{
		InputedUsers: []*pkg.ResponseUser{
			{
				DisplayName: "modaniru",
				ImageLink: "test",
			},
		},
		GeneralStreamers: []*pkg.Streamer{},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil{
		log.Fatal(err.Error())
	}
	
	s := grpc.NewServer()
	reflection.Register(s)
	pkg.RegisterTwitchGeneralFollowsServer(s, &MyServer{})

	if err = s.Serve(lis); err != nil{
		log.Fatal(err.Error())
	}
}