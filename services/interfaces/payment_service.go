package service_interfaces

import (
	"personal-growth/data/requests"

	"github.com/gofiber/fiber/v2"
)

type PaymentService interface {
	CreateMoMoPayment(request requests.PaymentRequest) (string, *fiber.Error)
	CreateVNPayPayment(request requests.PaymentRequest) (string, *fiber.Error)
	// SaveMoMoTransaction(request requests.MoMoTransactionRequest) (string, *fiber.Error)
	SaveVNPayTransaction(data requests.PaymentResultRequest) *fiber.Error
}
