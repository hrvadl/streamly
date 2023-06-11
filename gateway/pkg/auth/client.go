package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hrvadl/studdy-buddy/gateway/adapter"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/auth/pb"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
	"google.golang.org/grpc"
)

type AuthServiceClient struct {
	client pb.AuthClient
}

func (c *AuthServiceClient) HandleSignIn() gin.HandlerFunc {
	return adapter.Wrap[pb.SignInRequest, pb.SignInResponse](c.client.SignIn, adapter.WithBodyExtractor[pb.SignInRequest])
}

func NewClient(ac pb.AuthClient) *AuthServiceClient {
	return &AuthServiceClient{ac}
}

func InitClient(c *config.Config) pb.AuthClient {
	conn, err := grpc.Dial(c.AuthServiceURL, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect to auth service on URL %v %v", c.AuthServiceURL, err)
	}

	return pb.NewAuthClient(conn)
}