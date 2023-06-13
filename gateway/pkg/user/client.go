package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/adapter"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	client pb.UsersClient
}

func getUserIDFromURL(ctx *gin.Context) (*pb.GetByIdRequest, error) {
	return &pb.GetByIdRequest{
		Id: ctx.Param("id"),
	}, nil
}

func (c *UserServiceClient) HandleGetByID() gin.HandlerFunc {
	return adapter.Wrap[pb.GetByIdRequest, pb.GetByIdResponse](
		c.client.GetById, getUserIDFromURL,
	)
}

func NewService(uc pb.UsersClient) *UserServiceClient {
	return &UserServiceClient{uc}
}

func InitClient(c *config.Config) pb.UsersClient {
	conn, err := grpc.Dial(c.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("cannot connect to user service on URL %v %v", c.UserServiceURL, err)
	}

	return pb.NewUsersClient(conn)
}
