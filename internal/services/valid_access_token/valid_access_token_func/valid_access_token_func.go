package validaccesstokenfunc

import (
	"context"
	"errors"

	jwtconfig "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/config/jwt"
	pb "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/services/valid_access_token/valid_access_token_gen"
	"github.com/dgrijalva/jwt-go"
)

type Server struct {
	pb.UnimplementedValidAccessTokenServiceServer
}

func (s *Server) ValidAccessToken(ctx context.Context, req *pb.ValidRequest) (*pb.ValidResponse, error) {
	tokenString := req.AccessToken

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtconfig.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		return &pb.ValidResponse{
			Code: int32(401),
			Id:   int32(0),
			Role: "",
		}, errors.New("not auth")
	}

	claim := token.Claims.(jwt.MapClaims)
	return &pb.ValidResponse{
		Code: int32(200),
		Id:   int32(claim["id"].(float64)),
		Role: claim["role"].(string),
	}, nil
}
