package controllers

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/helpers"
	"personal-growth/models"
	service_interfaces "personal-growth/services/interfaces"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type AuthController struct {
	service service_interfaces.AuthService
}

func NewAuthController(service service_interfaces.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// @Summary      Login
// @Description  Login into the system
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body requests.LoginRequest true "Account info"
// @Success      200 {object} responses.LoginResponse
// @Router       /api/auth/login [POST]
func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	loginRequest := requests.LoginRequest{}
	err := ctx.BodyParser(&loginRequest)
	helpers.ErrorPanic(err)

	tokens, lerr := controller.service.Login(loginRequest)
	if lerr != nil {
		return ctx.Status(lerr.Code).JSON(lerr.Message)
	}

	// Set new refresh token in secure HttpOnly cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Login successfully",
		Data:    tokens,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// @Summary      Get new access token
// @Description  Get new access token using refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200 {object} string
// @Router       /api/auth/refresh [GET]
func (controller *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	// Lấy refresh_token từ cookie
	refreshToken := ctx.Cookies("refresh_token")
	if refreshToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No refresh token provided",
		})
	}

	tokens, lerr := controller.service.RefreshAccessToken(refreshToken)
	if lerr != nil {
		return ctx.Status(lerr.Code).JSON(lerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Get new acceess token successfully",
		Data:    tokens,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// @Summary      Register
// @Description  Register new account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body requests.RegisterRequest false "Registration info"
// @Success      200 {object} responses.UserResponse
// @Router       /api/auth/register [POST]
func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	request := requests.RegisterRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	user, rerr := controller.service.Register(request)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	var dto responses.UserResponse
	copier.Copy(&dto, user)
	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "You have registered a new account. Please check your mailbox to verify.",
		Data:    dto,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Me
// @Description  Get user data
// @Tags         Auth
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Success      200 {object} responses.UserResponse
// @Router       /api/auth/me [GET]
func (controller *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")

	var dto responses.UserResponse
	copier.Copy(&dto, user)
	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Get account data successfully",
		Data:    dto,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Change password
// @Description  Change user password
// @Tags         Auth
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        password body requests.ChangePasswordRequest true "Password info"
// @Router       /api/auth/change-password [POST]
func (controller *AuthController) ChangePassword(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)

	request := requests.ChangePasswordRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	rerr := controller.service.ChangePassword(request, user)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Change password successfully",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Forgot password
// @Description  Forgot user password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        password body requests.ForgotPasswordRequest true "Password info"
// @Router       /api/auth/forgot-password [POST]
func (controller *AuthController) ForgotPassword(ctx *fiber.Ctx) error {
	request := requests.ForgotPasswordRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	rerr := controller.service.ForgotPassword(request.Email)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Sent an otp to your email",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Verify OTP
// @Description  Verify OTP
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        otp body requests.VerifyOTPRequest true "OTP info"
// @Router       /api/auth/verify-otp [POST]
func (controller *AuthController) VerifyOTP(ctx *fiber.Ctx) error {
	request := requests.VerifyOTPRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	rerr := controller.service.VerifyOtp(request)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Verify your OTP successfully",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Verify Account
// @Description  Verify Account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        otp body requests.VerifyOTPRequest true "OTP info"
// @Router       /api/auth/verify-account [POST]
func (controller *AuthController) VerifyAccount(ctx *fiber.Ctx) error {
	request := requests.VerifyOTPRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	rerr := controller.service.VerifyAccount(request)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Verify your account successfully",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Resend OTP
// @Description  Resend OTP
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        otp body requests.ResendOTPRequest true "OTP info"
// @Router       /api/auth/resend-otp [POST]
func (controller *AuthController) ResendOTP(ctx *fiber.Ctx) error {
	request := requests.ResendOTPRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	rerr := controller.service.ResendOtp(request.Email)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Resend OTP to your email successfully",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Set new password
// @Description  Set new password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        password body requests.SetNewPasswordRequest true "Password info"
// @Router       /api/auth/set-new-password [POST]
func (controller *AuthController) SetNewPassword(ctx *fiber.Ctx) error {
	request := requests.SetNewPasswordRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	rerr := controller.service.SetNewPassword(request)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Set new password successfully",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary      Upload avatar
// @Description  Upload user avatar
// @Tags         Auth
// @Security 	 BearerAuth
// @Produce      json
// @Accept 		 multipart/form-data
// @Param 		 file formData file true "File to upload"
// @Router       /api/auth/upload-avatar [POST]
func (controller *AuthController) UploadAvatar(ctx *fiber.Ctx) error {
	file := ctx.Locals("file").(string)
	user := ctx.Locals("user").(*models.User)

	rerr := controller.service.UploadAvatar(file, user)
	if rerr != nil {
		return ctx.Status(rerr.Code).JSON(rerr.Message)
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Upload avatar successfully",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
