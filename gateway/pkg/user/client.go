package user

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
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
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return nil, err
	}

	return &pb.GetByIdRequest{
		Id: int32(id),
	}, nil
}

func (c *UserServiceClient) HandleGetByID() gin.HandlerFunc {
	return adapter.Wrap[pb.GetByIdRequest, pb.GetByIdResponse](
		c.client.GetById, getUserIDFromURL,
	)
}

func (c *UserServiceClient) HandleCreate() gin.HandlerFunc {
	return adapter.Wrap[pb.CreateRequest, pb.CreateResponse](
		c.client.Create, adapter.WithBodyExtractor[pb.CreateRequest],
	)
}

func (c *UserServiceClient) HandleChangePassword() gin.HandlerFunc {
	return adapter.Wrap[pb.ResetPasswordRequest, empty.Empty](
		c.client.ResetPassword, adapter.WithBodyExtractor[pb.ResetPasswordRequest],
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
