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
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		payemnt body requests.MoMoRequest true "Payment Info"
// @Router 		/api/payment/momo [post]
func (controller *PaymentController) MakeMomoPayment(ctx *fiber.Ctx) error {
	request := requests.MoMoRequest{}
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

// @Summary 	Get payment result
// @Description Get payment result
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Router 		/api/payment/momoreturn [post]
func (controller *PaymentController) MoMoReturnPayment(ctx *fiber.Ctx) error {
	println(ctx.Response())
	return ctx.Status(200).JSON("")
}

// @Summary 	Get payment notification
// @Description Get payment notification
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Router 		/api/payment/momonotify [get]
func (controller *PaymentController) MoMoNotifyPayment(ctx *fiber.Ctx) error {
	println(ctx.Response())
	return ctx.Status(200).JSON("")
}
