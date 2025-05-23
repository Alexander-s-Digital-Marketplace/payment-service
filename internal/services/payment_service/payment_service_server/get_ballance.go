package paymentserviceserver

import (
	"context"
	"errors"

	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models"
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func (s *Server) GetBallance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	var code int
	wallet := models.Wallet{
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

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/ZSDV7NU0M9XvdB3-DqGfaQjfbw9-Bi0s")
	if err != nil {
		logrus.Errorln(err)
		return &pb.GetBalanceResponse{
			Code:    int32(code),
			Balance: 0,
			Message: "Error connect to blockchain",
		}, errors.New("error connect to blockchain")
	}

	address := common.HexToAddress(wallet.WalletAddress)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		logrus.Errorln(err)
		return &pb.GetBalanceResponse{
			Code:    int32(code),
			Balance: 0,
			Message: "Error get balance",
		}, errors.New("error get balance")
	}

	return &pb.GetBalanceResponse{
		Code:    int32(code),
		Balance: float64(balance.Int64()) / 1e18,
		Message: "Success get balance",
	}, nil
}
