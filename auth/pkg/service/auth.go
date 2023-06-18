package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hrvadl/studdy-buddy/auth/pkg/config"
	"github.com/hrvadl/studdy-buddy/auth/pkg/models"
	"github.com/hrvadl/studdy-buddy/auth/pkg/pb"
	"github.com/hrvadl/studdy-buddy/auth/pkg/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	userReader *repositories.UserReader
	userWriter *repositories.UserWriter
	cfg        *config.Config
}

func New(ur *repositories.UserReader, uw *repositories.UserWriter, cfg *config.Config) *AuthService {
	return &AuthService{userReader: ur, userWriter: uw, cfg: cfg}
}

func (s *AuthService) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	profile, err := s.userReader.FindByEmail(in.Email)

	if err != nil {
		return nil, err
	}

	hash, err := s.GenerateHash(ctx, &pb.HashRequest{Str: in.Password})

	if err != nil {
		return nil, err
	}

	if profile.Password != hash.Hashed {
		return nil, status.Error(codes.InvalidArgument, "password is not the same")
	}

	token, err := s.GenerateToken(profile.Email, map[string]any{"id": profile.ID})

	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{
		Email: profile.Email,
		Login: profile.Email,
		Token: token,
	}, nil
}

func (s *AuthService) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	_, err := s.userReader.FindByEmailOrLogin(in.Email, in.Login)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.AlreadyExists, "User with such email already exists")
	}

	hash, err := s.GenerateHash(ctx, &pb.HashRequest{Str: in.Password})

	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong")
	}

	user := &models.User{
		Login:    in.Login,
		Password: hash.Hashed,
		Email:    in.Email,
	}

	if _, err := s.userWriter.Create(user); err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong with creating user")
	}

	return &pb.SignUpResponse{}, nil
}

func (s *AuthService) GenerateHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	hash := md5.Sum([]byte(in.Str))
	return &pb.HashResponse{Hashed: hex.EncodeToString(hash[:])}, nil
}

type TokenClaims struct {
	CustomClaims any `json:"customClaims"`
	jwt.RegisteredClaims
}

func (s *AuthService) GenerateToken(subject string, payload map[string]any) (string, error) {
	claims := TokenClaims{
		payload,
		jwt.RegisteredClaims{
			Subject:   subject,
			Issuer:    s.cfg.TokenIssuer,
			Audience:  jwt.ClaimStrings{s.cfg.TokenAudience},
			IssuedAt:  &jwt.NumericDate{time.Now()},
			ExpiresAt: &jwt.NumericDate{time.Now()},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(s.cfg.JwtKey))
}
