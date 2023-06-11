package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hrvadl/studdy-buddy/gateway/adapter"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/config"
	"github.com/hrvadl/studdy-buddy/gateway/pkg/user/pb"
	"google.golang.org/grpc"
)

type UserServiceClient struct {
	client pb.UsersClient
}

func (c *UserServiceClient) HandleGetByID() gin.HandlerFunc {
	return adapter.Wrap[pb.GetByIdRequest, pb.GetByIdResponse](
		c.client.GetById,
		func(ctx *gin.Context) (*pb.GetByIdRequest, error) {
			return &pb.GetByIdRequest{
				Id: ctx.Param("id"),
			}, nil
		},
	)
}

func NewClient(uc pb.UsersClient) *UserServiceClient {
	return &UserServiceClient{uc}
}

func InitClient(c *config.Config) pb.UsersClient {
	conn, err := grpc.Dial(c.UserServiceURL, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect to user service on URL %v %v", c.UserServiceURL, err)
	}

	return pb.NewUsersClient(conn)
}