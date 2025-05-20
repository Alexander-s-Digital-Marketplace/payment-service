package validaccesstokenfuncclient

import (
	"context"
	"time"

	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/valid_access_token/valid_access_token_gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ValidAccessToken(accessToken string) (int, int, string) {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, 0, ""
	}
	defer conn.Close()

	client := pb.NewValidAccessTokenServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ValidRequest{AccessToken: accessToken}

	res, err := client.ValidAccessToken(ctx, req)
	if err != nil {
		logrus.Errorln("Error send message:", err)
		return 503, 0, ""
	}

	return int(res.Code), int(res.Id), string(res.Role)
}
