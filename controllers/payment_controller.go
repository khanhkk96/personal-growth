package controllers

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/helpers"
	service_interfaces "personal-growth/services/interfaces"

	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	service service_interfaces.PaymentService
}

func NewPaymentController(service service_interfaces.PaymentService) *PaymentController {
	return &PaymentController{
		service: service,
	}
}

// @Summary 	Make a payment via MoMo
// @Description Make a payment via MoMo
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Param 		payemnt body requests.PaymentRequest true "Payment Info"
// @Success      200 {object} responses.Response
// @Router 		/api/payment/momo [post]
func (controller *PaymentController) MakeMomoPayment(ctx *fiber.Ctx) error {
	request := requests.PaymentRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	url, cerr := controller.service.CreateMoMoPayment(request)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Pay via momo successfully",
		Data:    url,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Get momo payment result
// @Description Get momo payment result
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Success      200 {object} responses.Response
// @Router 		/api/payment/momo_return [get]
func (controller *PaymentController) MoMoReturnPayment(ctx *fiber.Ctx) error {
	request := requests.MomoPaymentResultRequest{}
	if err := ctx.QueryParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := controller.service.SaveMomoTransaction(request); err != nil {
		return ctx.Status(err.Code).JSON(err.Message)
	}

	return ctx.Status(200).JSON("OK")
}

// @Summary 	Get momo payment notification
// @Description Get momo payment notification
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Success      200 {object} responses.Response
// @Router 		/api/payment/momo_notify [get]
func (controller *PaymentController) MoMoNotifyPayment(ctx *fiber.Ctx) error {
	println("MoMoNotifyPayment:::::")
	println(ctx.Request().URI().QueryArgs().String())

	return ctx.Status(200).JSON("")
}

// @Summary 	Make a payment via VNPAY
// @Description Make a payment via VNPAY
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Param 		payemnt body requests.PaymentRequest true "Payment Info"
// @Success      200 {object} responses.Response
// @Router 		/api/payment/vnpay [post]
func (controller *PaymentController) MakeVNPayPayment(ctx *fiber.Ctx) error {
	request := requests.PaymentRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	url, cerr := controller.service.CreateVNPayPayment(request)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Pay via VNPAY successfully",
		Data:    url,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Get vnpay payment result
// @Description Get vnpay payment result
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Success      200 {object} responses.Response
// @Router 		/api/payment/vnpay_return [get]
func (controller *PaymentController) VnpayReturnPayment(ctx *fiber.Ctx) error {
	var request requests.VNPayPaymentResultRequest

	if err := ctx.QueryParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := controller.service.SaveVNPayTransaction(request); err != nil {
		return ctx.Status(err.Code).JSON(err.Message)
	}

	return ctx.Status(200).JSON("OK")
}

// @Summary 	Get transaction list
// @Description Get the list of the transactions
// @Tags 		Payment
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} responses.PaymentPageResponse
// @Param 		filters query requests.PaymentFilters false "Payment Filter"
// @Router 		/api/payment [get]
func (controller *PaymentController) GetTransactions(ctx *fiber.Ctx) error {
	var filters requests.PaymentFilters
	if err := ctx.QueryParser(&filters); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// fmt.Printf("query::::::: %v\n", filters)

	data := controller.service.List(filters)

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Get list of transactions successfully",
		Data:    data,
	}

	return ctx.Status(200).JSON(response)
}
