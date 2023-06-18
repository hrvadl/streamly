package repositories

import (
	"context"
	"log"

	"github.com/hrvadl/studdy-buddy/auth/pkg/config"
	"github.com/hrvadl/studdy-buddy/auth/pkg/models"
	"github.com/hrvadl/studdy-buddy/auth/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

type UserReader struct {
	db *gorm.DB
}

func NewUserReader(db *gorm.DB) *UserReader {
	return &UserReader{db: db}
}

func (u *UserReader) FindByEmailOrLogin(email, login string) (*models.User, error) {
	var user models.User
	res := u.db.Where("Email = ?", email).Or("Login = ?", login).First(&user)
	return &user, res.Error
}

func (u *UserReader) FindByEmail(email string) (*models.User, error) {
	var user models.User
	res := u.db.Where("Email = ?", email).First(&user)
	return &user, res.Error
}

type UserWriter struct {
	pb.UsersClient
}

func NewUserWriter(c *config.Config) *UserWriter {
	conn, err := grpc.Dial(c.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("cannot connect to user service on URL %v %v", c.UserServiceURL, err)
	}

	return &UserWriter{pb.NewUsersClient(conn)}
}

func (u *UserWriter) Create(usr *models.User) (*int32, error) {
	ctx := context.Background()
	res, err := u.UsersClient.Create(ctx, &pb.CreateRequest{
		Login:    usr.Login,
		Email:    usr.Email,
		Password: usr.Password,
	})

	if err != nil {
		return nil, err
	}

	return &res.Id, nil
}
