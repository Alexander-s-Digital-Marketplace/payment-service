package paymentserviceserver

import (
	"context"
	"errors"
	"time"

	models "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models"
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
	"github.com/sirupsen/logrus"
)

func (s *Server) BuyProduct(ctx context.Context, req *pb.BuyProductRequest) (*pb.BuyProductResponse, error) {
	var code int
	walletSeller := models.Wallet{
		Id: int(req.WalletIdSeller),
	}
	code = walletSeller.GetFromTableById()
	if code != 200 {
		return &pb.BuyProductResponse{
			Code: int32(code),
		}, errors.New("error get from table")
	}

	walletBuyer := models.Wallet{
		Id: int(req.WalletIdBuyer),
	}
	code = walletBuyer.GetFromTableById()
	if code != 200 {
		return &pb.BuyProductResponse{
			Code: int32(code),
		}, errors.New("error get from table")
	}

	order := models.Order{
		ContractAddress:     "0x7097449F14dE64590F38A0eAa7ce946DC96fdd3c",
		ProductId:           int(req.ProductId),
		SellerWalletAddress: walletSeller.WalletAddress,
		BuyerWalletAddress:  walletBuyer.WalletAddress,
		ProductPrice:        req.ProductPrice,
		DateCreateOrder:     time.Now().Format("2006-01-02 15:04"),
	}
	logrus.Infoln("================================")
	logrus.Infoln("req.ProductId: ", req.ProductId)
	logrus.Infoln(order)
	logrus.Infoln("================================")
	code = order.AddToTable()
	if code != 200 {
		return &pb.BuyProductResponse{
			Code: int32(code),
		}, errors.New("error create order")
	}

	return &pb.BuyProductResponse{
		Code:          int32(code),
		OrderId:       int32(order.Id),
		Address:       order.ContractAddress,
		SellerAddress: order.SellerWalletAddress,
		ProductPrice:  order.ProductPrice,
	}, nil
}
