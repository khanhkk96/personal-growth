package routers

import (
	"personal-growth/common/constants"
	"personal-growth/controllers"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewAuthRouter(controller *controllers.AuthController, db *gorm.DB) *fiber.App {
	authRouter := fiber.New()

	authRouter.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controller.Register)
		router.Post("/login", controller.Login)
		router.Get("/refresh", controller.RefreshToken)
		router.Post("/forgot-password", controller.ForgotPassword)
		router.Post("/resend-otp", controller.ResendOTP)
		router.Post("/verify-otp", controller.VerifyOTP)
		router.Post("/verify-account", controller.VerifyAccount)
		router.Post("/set-new-password", controller.SetNewPassword)
	})

	authRouter.Group("/auth", middlewares.Authenticate(), middlewares.GetProfileHandler(db)).Route("/", func(router fiber.Router) {
		router.Get("/me", controller.Me)
		router.Get("/logout", controller.Logout)
		router.Post("/change-password", controller.ChangePassword)
		router.Post("/upload-avatar", middlewares.UploadFileHandlder(middlewares.UploadFileOptions{
			AllowedTypes: constants.ImageFileTypes,
			Destination:  "avatar",
		}), controller.UploadAvatar)
	})

	return authRouter
}
