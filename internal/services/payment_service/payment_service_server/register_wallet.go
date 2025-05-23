package paymentserviceserver

import (
	"context"
	"errors"

	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models"
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
)

func (s *Server) RegisterWallet(ctx context.Context, req *pb.RegisterWalletRequest) (*pb.RegisterWalletResponse, error) {
	var code int
	wallet := models.Wallet{
		WalletAddress: req.WalletAddress,
	}

	code = wallet.AddToTable()
	if code != 200 {
		return &pb.RegisterWalletResponse{
			Code:    int32(code),
			Message: "Error add to table",
		}, errors.New("error add to table")
	}

	return &pb.RegisterWalletResponse{
		Code:     int32(code),
		Message:  "Success add to table",
		WalletId: int32(wallet.Id),
	}, nil

}
