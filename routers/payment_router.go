package routers

import (
	"personal-growth/controllers"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewPaymentRouter(controller *controllers.PaymentController, db *gorm.DB) *fiber.App {
	paymentRouter := fiber.New()

	requiredAuthRouter := paymentRouter.Group("/", middlewares.Authenticate(), middlewares.GetProfileHandler(db))

	requiredAuthRouter.Route("/payment", func(router fiber.Router) {
		router.Post("/momo", controller.MakeMomoPayment)
	})

	paymentRouter.Route("/payment", func(router fiber.Router) {
		router.Get("/momonotify", controller.MoMoReturnPayment)
		router.Get("/momoreturn", controller.MoMoNotifyPayment)
	})

	return paymentRouter
}
