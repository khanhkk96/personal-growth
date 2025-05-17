package services

import (
	"fmt"
	"personal-growth/data/requests"
	"personal-growth/helpers"
	"personal-growth/repositories"
	service_interfaces "personal-growth/services/interfaces"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PaymentServiceImpl struct {
	repository repositories.PaymentRepository
	validate   *validator.Validate
}

func NewPaymentServiceImpl(repository repositories.PaymentRepository, validate *validator.Validate) service_interfaces.PaymentService {
	return &PaymentServiceImpl{
		repository: repository,
		validate:   validate,
	}
}

// createMoMoPayment implements service_interfaces.PaymentService.
func (p *PaymentServiceImpl) CreateMoMoPayment(request requests.MoMoRequest) (string, *fiber.Error) {
	url, err := helpers.PayViaMoMo(request.Amount, request.Description)
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "Make a payment by MOMO unsuccessfully")
	}

	fmt.Println("Payment URL:", url)

	return url, nil
}
