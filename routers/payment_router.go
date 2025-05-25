package routers

import (
	"personal-growth/controllers"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewPaymentRouter(controller *controllers.PaymentController, db *gorm.DB) *fiber.App {
	paymentRouter := fiber.New()

	paymentRouter.Route("/payment", func(router fiber.Router) {
		router.Get("/momo_notify", controller.MoMoNotifyPayment)
		router.Get("/momo_return", controller.MoMoReturnPayment)
		router.Get("/vnpay_return", controller.VnpayReturnPayment)
		router.Post("/momo", controller.MakeMomoPayment)
		router.Post("/vnpay", controller.MakeVNPayPayment)
	})

	paymentRouter.Group("/payment", middlewares.Authenticate(), middlewares.GetProfileHandler(db)).Route("/", func(router fiber.Router) {
		// router.Post("/momo", controller.MakeMomoPayment)
		// router.Post("/vnpay", controller.MakeVNPayPayment)
		router.Get("/", controller.GetTransactions)
	})

	return paymentRouter
}
