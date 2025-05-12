package paymentserviceserver

import (
	"context"
	"errors"

	walletmodel "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models/wallet_model"
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
	"github.com/sirupsen/logrus"
)

func GetBallance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	var code int
	wallet := walletmodel.Wallet{
		Id: int(req.WalletId),
	}
	code = wallet.GetFromTableById()
	if code != 200 {
		return &pb.GetBalanceResponse{
			Code:    int32(code),
			Balance: 0,
			Message: "Error get from table",
		}, errors.New("error get from table")
	}

	//calling smart contract
	logrus.Infoln("Calling smart contract to get balance")

	return &pb.GetBalanceResponse{
		Code:    int32(code),
		Balance: 500,
		Message: "Success get balance",
	}, nil
}
