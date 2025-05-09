package router

import (
	"personal-growth/common/constants"
	"personal-growth/controller"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewAuthRouter(controller *controller.AuthController, db *gorm.DB) *fiber.App {
	appRouter := fiber.New()

	appRouter.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controller.Register)
		router.Post("/login", controller.Login)
		router.Get("/refresh", controller.RefreshToken)
		router.Post("/forgot-password", controller.ForgotPassword)
		router.Post("/resend-otp", controller.ResendOTP)
		router.Post("/verify-otp", controller.VerifyOTP)
		router.Post("/verify-account", controller.VerifyAccount)
		router.Post("/set-new-password", controller.SetNewPassword)
	})

	authRouter := appRouter.Group("/", middlewares.Authenticate(), middlewares.GetProfileHandler(db))

	// validate authentication middleware
	authRouter.Route("/auth", func(router fiber.Router) {
		router.Get("/me", controller.Me)
		router.Post("/change-password", controller.ChangePassword)
		router.Post("/upload-avatar", middlewares.Uploadfile(middlewares.UploadFileOptions{
			AAllowedTypes: constants.ImageFileTypes,
		}), controller.UploadAvatar)
	})

	return appRouter
}
