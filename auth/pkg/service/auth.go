package service

import (
	"context"

	"github.com/hrvadl/studdy-buddy/auth/pkg/pb"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func New() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return &pb.SignInResponse{}, nil
}

func (s *AuthService) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return &pb.SignUpResponse{}, nil
}
