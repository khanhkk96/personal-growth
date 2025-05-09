package service

import (
	"personal-growth/data/request"
	"personal-growth/data/response"
	"personal-growth/model"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(data request.LoginRequest) (*response.LoginResponse, *fiber.Error)
	RefreshAccessToken(refreshToken string) (string, *fiber.Error)
	Register(user request.RegisterRequest) (*model.User, *fiber.Error)
	ForgotPassword(email string) *fiber.Error
	VerifyOtp(data request.VerifyOTPRequest) *fiber.Error
	VerifyAccount(data request.VerifyOTPRequest) *fiber.Error
	ChangePassword(data request.ChangePasswordRequest, user *model.User) *fiber.Error
	ResendOtp(email string) *fiber.Error
	SetNewPassword(data request.SetNewPasswordRequest) *fiber.Error
}
