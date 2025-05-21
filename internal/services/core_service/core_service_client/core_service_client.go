package coreserviceclient

import (
	"context"
	"time"

	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/core_service/core_service_gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConfirmPaymentOfProduct(orderId int, productId int, walletId int) (int, string) {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, "Error connent to localhost:50052"
	}
	defer conn.Close()

	client := pb.NewCoreServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UpdateSoldProductRequest{
		OrderId:   int32(orderId),
		ProductId: int32(productId),
		WalletId:  int32(walletId),
	}

	res, err := client.UpdateSoldProduct(ctx, req)
	if err != nil {
		logrus.Errorln("Error send message:", err)
		return 503, res.Message
	}

	return int(res.Code), res.Message
}
