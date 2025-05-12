package paymentserviceserver

import (
	pb "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/payment_service/payment_service_gen"
)

type Server struct {
	pb.UnimplementedPaymentServiceServer
}
