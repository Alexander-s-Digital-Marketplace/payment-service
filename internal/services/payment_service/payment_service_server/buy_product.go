package paymentserviceserver

import (
	"context"
	"errors"

	walletmodel "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models/wallet_model"
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
	"github.com/sirupsen/logrus"
)

func BuyProduct(ctx context.Context, req *pb.BuyProductRequest) (*pb.BuyProductResponse, error) {
	var code int
	walletSeller := walletmodel.Wallet{
		Id: int(req.WalletIdSeller),
	}
	code = walletSeller.GetFromTableById()
	if code != 200 {
		return &pb.BuyProductResponse{
			Code:    int32(code),
			Message: "Error get from table",
		}, errors.New("error get from table")
	}

	walletBuyer := walletmodel.Wallet{
		Id: int(req.WalletIdBuyer),
	}
	code = walletBuyer.GetFromTableById()
	if code != 200 {
		return &pb.BuyProductResponse{
			Code:    int32(code),
			Message: "Error get from table",
		}, errors.New("error get from table")
	}

	//price := req.ProductPrice
	//calling smart contract
	logrus.Infoln("Calling smart contract to buy product")

	return &pb.BuyProductResponse{
		Code:    int32(code),
		Message: "Success buy product",
	}, nil
}
