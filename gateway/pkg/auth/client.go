package auth

import (
	"log"

	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
	"google.golang.org/grpc"
)

func NewServiceClient(c *config.Config) {
	conn, err := grpc.Dial(c.AuthServiceURL)

	if err != nil {
		log.Fatalf("cannot connect to auth service %v", err)
	}
	conn.Close()
}
