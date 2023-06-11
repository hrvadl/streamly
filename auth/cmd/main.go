package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hrvadl/studdy-buddy/auth/pkg/config"
	"github.com/hrvadl/studdy-buddy/auth/pkg/pb"
	"github.com/hrvadl/studdy-buddy/auth/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	c := config.Load()
	TCPServer, err := net.Listen("tcp", fmt.Sprintf(":%v", c.Port))

	if err != nil {
		log.Fatalf("cannot listen on TCP PORT %v %v", c.Port, err)
	}

	authService := service.New()

	gRPCServer := grpc.NewServer()
	pb.RegisterAuthServer(gRPCServer, authService)
	fmt.Printf("server listening on port %v", c.Port)
	gRPCServer.Serve(TCPServer)
}
