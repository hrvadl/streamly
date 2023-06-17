package service

import (
	"context"

	"github.com/hrvadl/studdy-buddy/auth/pkg/pb"
	"gorm.io/gorm"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	db *gorm.DB
}

func New(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{
		Email: "mail",
		Phone: "phone",
		Role:  "role",
	}, nil
}

func (s *AuthService) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return &pb.SignUpResponse{}, nil
}
