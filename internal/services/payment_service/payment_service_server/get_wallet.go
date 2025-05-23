package paymentserviceserver

import (
	"context"
	"errors"

	models "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models"
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
)

func (s *Server) GetWallet(ctx context.Context, req *pb.GetWalletRequest) (*pb.GetWalletResponse, error) {
	var code int
	wallet := models.Wallet{
		Id: int(req.WalletId),
	}
	code = wallet.GetFromTableById()
	if code != 200 {
		return &pb.GetWalletResponse{
			Code:    int32(code),
			Wallet:  "",
			Message: "Error get from table",
		}, errors.New("error get from table")
	}

	return &pb.GetWalletResponse{
		Code:    int32(code),
		Wallet:  wallet.WalletAddress,
		Message: "Success get wallet",
	}, nil
}
