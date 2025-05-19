package controllers

import (
	"net/url"
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/helpers"
	service_interfaces "personal-growth/services/interfaces"
	"personal-growth/utils"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
	println(ctx)
	println(ctx.Response())
	return ctx.Status(200).JSON("")
}

// @Summary 	Get momo payment notification
// @Description Get momo payment notification
// @Tags 		Payment
// @Accept 		json
// @Produce 	json
// @Success      200 {object} responses.Response
// @Router 		/api/payment/momo_notify [get]
func (controller *PaymentController) MoMoNotifyPayment(ctx *fiber.Ctx) error {
	println(ctx.Response())
	return ctx.Status(200).JSON("")
}

// @Summary 	Make a payment via VNPAY
// @Description Make a payment via VNPAY
// @Tags 		Payment
// @Security  	BearerAuth
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
	query := ctx.Request().URI().QueryArgs()
	vnpParams := make(map[string]string)

	query.VisitAll(func(k, v []byte) {
		key := string(k)
		val := string(v)
		if key != "vnp_SecureHash" && key != "vnp_SecureHashType" {
			vnpParams[key] = val
		}
	})

	// Tạo chuỗi hash
	keys := make([]string, 0, len(vnpParams))
	for k := range vnpParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var hashData strings.Builder
	for i, k := range keys {
		if i > 0 {
			hashData.WriteString("&")
		}
		hashData.WriteString(k + "=" + url.QueryEscape(vnpParams[k]))
	}

	println(hashData.String())

	secureHash := string(query.Peek("vnp_SecureHash"))
	vnp_HashSecret := viper.GetString("VNP_HASHSECRET")
	myHash := utils.CreateVNPayHash(vnp_HashSecret, hashData.String())
	println(secureHash)
	println(myHash)

	return ctx.Status(200).JSON("OK")
}
