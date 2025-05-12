package paymentserviceserver

import (
	"context"
	"errors"
	"net/http"

	walletmodel "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models/wallet_model"
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
	"github.com/sirupsen/logrus"
)

func UpdateWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.UpdateWalletResponse, error) {
	var code int
	walletOld := walletmodel.Wallet{
		Id: int(req.OldWalletId),
	}
	code = walletOld.GetFromTableById()
	if code != 200 {
		return &pb.UpdateWalletResponse{
			Code:    int32(code),
			Message: "Error get from table",
		}, errors.New("error get from table")
	}

	walletNew := walletmodel.Wallet{
		WalletAddress: req.NewWalletAddress,
	}
	if walletNew.WalletAddress == walletOld.WalletAddress {
		return &pb.UpdateWalletResponse{
			Code:    int32(http.StatusBadRequest),
			Message: "Wallet address is the same",
		}, errors.New("wallet address is the same")
	}
	code = walletNew.AddToTable()
	if code != 200 {
		return &pb.UpdateWalletResponse{
			Code:    int32(code),
			Message: "Error add to table",
		}, errors.New("error add to table")
	}
	code = walletOld.DeleteFromTable()
	if code != 200 {
		return &pb.UpdateWalletResponse{
			Code:    int32(code),
			Message: "Error delete from table",
		}, errors.New("error delete from table")
	}

	//calling smart contract
	logrus.Infoln("Calling smart contract to update wallet")

	return &pb.UpdateWalletResponse{
		Code:    int32(code),
		Message: "Success update wallet",
	}, nil
}
