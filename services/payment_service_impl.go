package services

import (
	"fmt"
	"log"
	"personal-growth/data/requests"
	"personal-growth/helpers"
	"personal-growth/models"
	"personal-growth/repositories"
	service_interfaces "personal-growth/services/interfaces"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
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
func (p *PaymentServiceImpl) CreateMoMoPayment(request requests.PaymentRequest) (string, *fiber.Error) {
	url, err := helpers.PayViaMoMoLink(request.Amount, request.Description)
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "Make a payment by MOMO unsuccessfully")
	}

	fmt.Println("Payment URL:", url)

	return url, nil
}

func (p *PaymentServiceImpl) CreateVNPayPayment(request requests.PaymentRequest) (string, *fiber.Error) {
	url, err := helpers.PayViaVNPay(request.Amount, request.Description)
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "Make a payment by VNPAY unsuccessfully")
	}

	fmt.Println("Payment URL:", url)

	return url, nil
}

func (p *PaymentServiceImpl) SaveVNPayTransaction(data requests.PaymentResultRequest) *fiber.Error {
	if data.TransactionStatus != "00" {
		return fiber.NewError(fiber.StatusBadRequest, "Transaction is not paid")
	}

	parsedTime, err := time.Parse("20060102150405", data.PayDate)
	helpers.ErrorPanic(err)

	payment := &models.Payment{}
	copier.Copy(payment, data)
	payment.TransactionStatus = "success"
	payment.PayDate = parsedTime
	payment.Amount = data.Amount / 100
	payment.PayBy = "vnpay"

	cerr := p.repository.Create(payment)
	if cerr != nil {
		log.Println("Error: ", cerr, data)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save new payment")
	}

	return nil
}
