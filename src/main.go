package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pkg "github.com/modaniru/tgf-gRPC/pkg/proto"
	"github.com/modaniru/tgf-gRPC/src/client"
	"github.com/modaniru/tgf-gRPC/src/server"
	"github.com/modaniru/tgf-gRPC/src/service"
	"github.com/modaniru/tgf-gRPC/src/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
// TODO add Makefile for building
// TODO add README.md
func main() {
	//Load yaml and .env file
	err := utils.LoadConfig("configuration/", "yaml")
	if err != nil{
		log.Fatal(err.Error())
	}
	//DIP
	client := client.NewQueries(os.Getenv("TWITCH_CLIENT_ID"), os.Getenv("TWITCH_CLIENT_SECRET"))
	service := service.NewService(client)
	server := server.NewServer(service)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("port")))
	if err != nil{
		log.Fatal(err.Error())
	}
	//gRPC server
	s := grpc.NewServer()
	reflection.Register(s)
	pkg.RegisterTwitchGeneralFollowsServer(s, server)

	if err = s.Serve(lis); err != nil{
		log.Fatal(err.Error())
	}
}