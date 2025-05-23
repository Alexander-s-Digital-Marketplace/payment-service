package paymentlistener

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"os"
	"time"

	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models"
	coreserviceclient "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/core_service/core_service_client"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func ListenForPayment() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/ZSDV7NU0M9XvdB3-DqGfaQjfbw9-Bi0s")
	if err != nil {
		logrus.Fatalf("ethclient: %v", err)
	}
	contractAddr := common.HexToAddress("0x7097449F14dE64590F38A0eAa7ce946DC96fdd3c")

	// Чтение ABI из json-артефакта
	type Artifact struct {
		ABI json.RawMessage `json:"abi"`
	}
	var artifact Artifact
	data, err := os.ReadFile("internal/services/payment_service/payment_listener/PayRouter.json")
	if err != nil {
		logrus.Fatal(err)
	}
	if err := json.Unmarshal(data, &artifact); err != nil {
		logrus.Fatalf("Unmarshal artifact: %v", err)
	}
	contractAbi, err := abi.JSON(bytes.NewReader(artifact.ABI))
	if err != nil {
		logrus.Fatalf("abi: %v", err)
	}

	// Получаем сигнатуру события (только для фильтрации по конкретному событию)
	eventSignature := []byte("ProductPaid(uint256,address,address,uint256)")
	eventSigHash := crypto.Keccak256Hash(eventSignature)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		Topics:    [][]common.Hash{{eventSigHash}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logrus.Fatalf("subscribe: %v", err)
	}

	eventName := "ProductPaid"
	logrus.Infoln("Слушаем события ProductPaid...")

	for {
		select {
		case err := <-sub.Err():
			logrus.Errorln("Ошибка подписки:", err)
			return
		case vLog := <-logs:
			var event struct {
				OrderId *big.Int
				Buyer   common.Address
				Seller  common.Address
				Amount  *big.Int
			}
			err := contractAbi.UnpackIntoInterface(&event, eventName, vLog.Data)
			if err != nil {
				logrus.Errorln("unpack:", err)
				continue
			}

			// Индексированные параметры извлекаются из topics
			event.Buyer = common.HexToAddress(vLog.Topics[1].Hex())
			event.Seller = common.HexToAddress(vLog.Topics[2].Hex())

			logrus.Infof("Оплата! OrderID=%v, Buyer=%s, Seller=%s, amount=%s wei, tx=%s\n",
				event.OrderId, event.Buyer.Hex(), event.Seller.Hex(), event.Amount.String(), vLog.TxHash.Hex(),
			)

			orderOld := models.Order{
				Id: int(event.OrderId.Int64()),
			}

			code := orderOld.GetFromTableById()
			if code != 200 {
				logrus.Errorln("Error get order from table")
			}

			order := models.Order{
				Id:                  int(event.OrderId.Int64()),
				ContractAddress:     contractAddr.Hex(),
				SellerWalletAddress: event.Seller.Hex(),
				BuyerWalletAddress:  event.Buyer.Hex(),
				ProductPrice:        float64(event.Amount.Int64()) / 1e18,
				TxHash:              vLog.TxHash.Hex(),
				IsPaid:              true,
				DatePaidOrder:       time.Now().Format("2006-01-02 15:04"),
				DateCreateOrder:     orderOld.DateCreateOrder,
				ProductId:           orderOld.ProductId,
			}
			code = order.UpdateInTable()
			if code != 200 {
				logrus.Errorln("Error update order in table")
			}

			walletBuyer := models.Wallet{
				WalletAddress: order.BuyerWalletAddress,
			}
			code = walletBuyer.GetFromTableByAddress()
			if code != 200 {
				logrus.Errorln("Error get walletBuyer from table")
			}

			var errr string
			code, errr = coreserviceclient.ConfirmPaymentOfProduct(order.Id, order.ProductId, walletBuyer.Id)
			if code != 200 {
				logrus.Errorln(errr)
			}

			logrus.Infoln("Success update order in table")
		}
	}
}
