package service_interfaces

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"

	"github.com/gofiber/fiber/v2"
)

type PaymentService interface {
	CreateMoMoPayment(request requests.PaymentRequest) (string, *fiber.Error)
	CreateVNPayPayment(request requests.PaymentRequest) (string, *fiber.Error)
	SaveMomoTransaction(data requests.MomoPaymentResultRequest) *fiber.Error
	SaveVNPayTransaction(data requests.VNPayPaymentResultRequest) *fiber.Error
	List(filters requests.PaymentFilters) responses.PaymentPageResponse
}
