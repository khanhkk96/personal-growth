package service_interfaces

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/db/models"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(data requests.LoginRequest) (*responses.LoginResponse, *fiber.Error)
	RefreshAccessToken(refreshToken string) (string, *fiber.Error)
	Register(user requests.RegisterRequest) (*models.User, *fiber.Error)
	ForgotPassword(email string) *fiber.Error
	VerifyOtp(data requests.VerifyOTPRequest) *fiber.Error
	VerifyAccount(data requests.VerifyOTPRequest) *fiber.Error
	ChangePassword(data requests.ChangePasswordRequest, user *models.User) *fiber.Error
	ResendOtp(email string) *fiber.Error
	SetNewPassword(data requests.SetNewPasswordRequest) *fiber.Error
	UploadAvatar(file string, user *models.User) *fiber.Error
	Logout(uid string, refreshToken string) *fiber.Error
}
