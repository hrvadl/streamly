package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hrvadl/studdy-buddy/auth/pkg/config"
	"github.com/hrvadl/studdy-buddy/auth/pkg/db"
	"github.com/hrvadl/studdy-buddy/auth/pkg/pb"
	"github.com/hrvadl/studdy-buddy/auth/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	l := log.Default()
	l.Print("auth service starting...")
	c := config.Load()
	db := db.Init(c)
	TCPServer, err := net.Listen("tcp", fmt.Sprintf(":%v", c.Port))

	if err != nil {
		l.Fatalf("cannot listen on TCP PORT %v %v", c.Port, err)
	}

	authService := service.New(db)

	gRPCServer := grpc.NewServer()
	pb.RegisterAuthServer(gRPCServer, authService)
	l.Printf("server listening on port %v", c.Port)
	gRPCServer.Serve(TCPServer)
}
