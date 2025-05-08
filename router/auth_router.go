package router

import (
	"personal-growth/controller"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewAuthRouter(controller *controller.AuthController, db *gorm.DB) *fiber.App {
	appRouter := fiber.New()

	appRouter.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controller.Register)   //done
		router.Post("/login", controller.Login)         //done
		router.Get("/refresh", controller.RefreshToken) //done
		router.Post("/forgot-password", controller.ForgotPassword)
		router.Post("/resend-otp", controller.ResendOTP)         //done
		router.Post("/verify-otp", controller.VerifyOTP)         //done
		router.Post("/verify-account", controller.VerifyAccount) //done
	})

	authRouter := appRouter.Group("/", middlewares.Authenticate(), middlewares.GetProfileHandler(db))

	// validate authentication middleware
	authRouter.Route("/auth", func(router fiber.Router) {
		router.Get("/me", controller.Me)                           //done
		router.Post("/change-password", controller.ChangePassword) //done
	})

	return appRouter
}
