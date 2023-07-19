package main

import (
	"log"
	"net"
	"os"

	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
	"github.com/modaniru/tgf-gRPC/src/client"
	"github.com/modaniru/tgf-gRPC/src/server"
	"github.com/modaniru/tgf-gRPC/src/service"
	"github.com/modaniru/tgf-gRPC/src/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO add README.md with docs
// case сикретов в образе
func main() {
	//Load .env file
	utils.LoadEnvIfExist()
	//DIP
	client := client.NewQueries(os.Getenv("TWITCH_CLIENT_ID"), os.Getenv("TWITCH_CLIENT_SECRET"))
	service := service.NewService(client)
	server := server.NewServer(service)
	lis, err := net.Listen("tcp", ":"+utils.GetPort())
	if err != nil {
		log.Fatal(err.Error())
	}
	//gRPC server
	s := grpc.NewServer()
	reflection.Register(s)
	pkg.RegisterTwitchGeneralFollowsServer(s, server)
	log.Println("server start. Port: ", utils.GetPort())
	if err = s.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
