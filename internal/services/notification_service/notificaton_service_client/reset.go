package notificationservice

import (
	"context"
	"time"

	resetpasswordmodel "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/reset_password_model"
	pb "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/services/notification_service/notification_service_gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ResetNotif(resetCode resetpasswordmodel.ResetCode, email string) (int, string) {
	conn, err := grpc.NewClient(
		"localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, ""
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ResetRequest{
		Email:     email,
		ResetCode: int32(resetCode.Code),
	}

	logrus.Infoln("req:", req)

	res, err := client.ResetNotif(ctx, req)
	if err != nil {
		logrus.Errorln("Error send message:", err)
		return 503, ""
	}

	return int(res.Code), res.Message
}
