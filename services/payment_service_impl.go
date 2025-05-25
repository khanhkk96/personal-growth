package services

import (
	"fmt"
	"log"
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/db/models"
	"personal-growth/helpers"
	"personal-growth/repositories"
	service_interfaces "personal-growth/services/interfaces"
	"personal-growth/utils"
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
	err := p.validate.Struct(request)
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, helpers.PrintErrorMessage(err))
	}

	url, err := helpers.PayViaQRMoMo(request.Amount, request.Description)
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "Make a payment by MOMO unsuccessfully")
	}

	fmt.Println("Payment URL:", url)

	return url, nil
}

func (p *PaymentServiceImpl) CreateVNPayPayment(request requests.PaymentRequest) (string, *fiber.Error) {
	err := p.validate.Struct(request)
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, helpers.PrintErrorMessage(err))
	}

	url, err := helpers.PayViaVNPay(request.Amount, request.Description)
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "Make a payment by VNPAY unsuccessfully")
	}

	fmt.Println("Payment URL:", url)

	return url, nil
}

func (p *PaymentServiceImpl) SaveVNPayTransaction(data requests.VNPayPaymentResultRequest) *fiber.Error {
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
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save new vnpay payment")
	}

	return nil
}

func (p *PaymentServiceImpl) SaveMomoTransaction(data requests.MomoPaymentResultRequest) *fiber.Error {
	if data.ResponseCode != "0" {
		return fiber.NewError(fiber.StatusBadRequest, "Transaction is not paid")
	}

	parsedTime, err := time.Parse("20060102150405", data.PayDate)
	helpers.ErrorPanic(err)

	payment := &models.Payment{}
	copier.Copy(payment, data)
	payment.TransactionStatus = "success"
	payment.PayDate = parsedTime
	payment.Amount = data.Amount / 100
	payment.PayBy = "momo"

	cerr := p.repository.Create(payment)
	if cerr != nil {
		log.Println("Error: ", cerr, data)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save new momo payment")
	}

	return nil
}

func (p *PaymentServiceImpl) List(options requests.PaymentFilters) responses.PaymentPageResponse {
	var transactions []models.Payment
	builder := p.repository.GetDataSource().Model(&models.Payment{})

	if !utils.IsEmpty(&options.Query) {
		queryByName := fmt.Sprintf(`%%%s%%`, options.Query)
		builder = builder.Where("txn_ref LIKE ? OR transaction_no LIKE ?", queryByName, queryByName)
	}

	if !utils.IsEmpty(&options.PayBy) {
		builder = builder.Where("pay_by = ?", options.PayBy)
	}

	if options.PayFrom != nil {
		builder = builder.Where("pay_date >= ?", *options.PayFrom)
	}

	if options.PayTo != nil {
		builder = builder.Where("pay_date <= ?", *options.PayTo)
	}

	var totalItem int64
	builder.Count(&totalItem)

	builder.Offset((options.Page - 1) * options.Limit).Limit(options.Limit).Order(fmt.Sprintf("%s %s", options.OrderBy, options.Order)).Find(&transactions)

	// Convert projects to []interface{}
	paymentResponses := make([]responses.PaymentResponse, len(transactions))
	for i, tran := range transactions {
		copier.Copy(&paymentResponses[i], tran.Model)
		copier.Copy(&paymentResponses[i], tran)
	}

	metadata := responses.NewPaginationMetaData(options.Page, options.Limit, int(totalItem), paymentResponses)

	return responses.NewPaginatedResponse(metadata)
}
