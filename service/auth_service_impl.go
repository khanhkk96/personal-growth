package service

import (
	"database/sql"
	"fmt"
	"personal-growth/config"
	"personal-growth/data/request"
	"personal-growth/data/response"
	"personal-growth/helpers"
	"personal-growth/model"
	"personal-growth/repository"
	"personal-growth/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.BaseRepository[model.User]
	validate       *validator.Validate
}

func NewAuthServiceImpl(repository repository.BaseRepository[model.User], validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: repository,
		validate:       validate,
	}
}

func (n *AuthServiceImpl) Login(data request.LoginRequest) (*response.LoginResponse, *fiber.Error) {
	// Validate username and password
	if err := n.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	// Check if user exists in the database
	user, err := n.UserRepository.FindOneBy("username = ? AND is_active = ?", data.Username, true)
	if err != nil || user == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	// Check if password is correct
	if !user.CompareHashAndPassword(data.Password) {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Password is wrong")
	}

	// generate access token
	config, _ := config.LoadConfig(".")
	token, rf, err := utils.GenerateTokens(config.TokenExpiredIn, user.Id, config.TokenSecret, config.RefreshTokenSecret)
	helpers.ErrorPanic(err)

	return &response.LoginResponse{
		AccessToken:  token,
		RefreshToken: rf,
	}, nil
}

func (n *AuthServiceImpl) RefreshAccessToken(refreshToken string) (string, *fiber.Error) {
	config, _ := config.LoadConfig(".")
	// Kiểm tra refresh token hợp lệ
	claims, err := utils.ValidateRefreshToken(refreshToken, config.RefreshTokenSecret)
	if err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired refresh token")
	}

	userID := claims["sub"].(string)

	// Tạo access token mới
	newAccessToken, err := utils.GenerateAccessToken(config.TokenExpiredIn, userID, config.TokenSecret)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Could not create access token")
	}

	return newAccessToken, nil
}

func (n *AuthServiceImpl) Register(data request.RegisterRequest) (*model.User, *fiber.Error) {
	// Validate input data
	if err := n.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	// Check if user exists in the database
	user, _ := n.UserRepository.FindOneBy("Username = ?", data.Username)
	if user != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "User exists")
	}

	//save user data
	user = &model.User{}
	copier.Copy(user, data)

	//generate OTP
	otp, _ := utils.GenerateNumberOTP(6)

	user.Otp = sql.NullString{String: otp, Valid: true}
	user.OtpExpiredAt = sql.NullTime{Time: time.Now().Add(5 * time.Minute), Valid: true}
	user.OtpCounter++
	// isActive := true
	// user.IsActive = &isActive

	cerr := n.UserRepository.Create(user)
	if cerr != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Register your account unsuccessfully")
	}

	//send verification email
	serr := helpers.SendEmail(data.Email, "Verify your account", fmt.Sprintf("Welcome to KK Project <br /> You have created new account succeessfully <br /> Please verify your one by using bellow OTP: <br /> <h2>%s</h2>", otp))
	if serr != nil {
		panic(serr)
	}

	// Return user information
	return user, nil
}

func (n *AuthServiceImpl) ForgotPassword(email string) *fiber.Error {
	err := n.validate.Var(email, "required,email")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email format")
	}

	// Check if user exists in the database
	user, err := n.UserRepository.FindOneBy("email = ?", email)
	if err != nil || user == nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	//generate OTP
	otp, _ := utils.GenerateNumberOTP(6)
	//send verification email
	serr := helpers.SendEmail(email, "Email verification", fmt.Sprintf("Your OTP: <br /> <h2>%s</h2>", otp))
	if serr != nil {
		panic(serr)
	}

	n.UserRepository.Update(user)
	return nil
}

func (n *AuthServiceImpl) VerifyAccount(data request.VerifyOTPRequest) *fiber.Error {
	// Check if user exists in the database
	err := n.VerifyOtp(data)
	if err != nil {
		return err
	}

	user, _ := n.UserRepository.FindOneBy("email = ?", data.Email)

	// clear OTP and expired time
	user.Otp = sql.NullString{Valid: false}
	user.OtpExpiredAt = sql.NullTime{Valid: false}
	user.OtpCounter = 0
	user.IsActive = true
	n.UserRepository.Update(user)

	return nil
}

func (n *AuthServiceImpl) VerifyOtp(data request.VerifyOTPRequest) *fiber.Error {
	// Validate input data
	if err := n.validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	// Check if user exists in the database
	user, err := n.UserRepository.FindOneBy("email = ?", data.Email)
	fmt.Println(user)
	if err != nil || user == nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	// Check if OTP is correct
	if user.Otp.String != data.Otp {
		return fiber.NewError(fiber.StatusBadRequest, "Incorrect OTP")
	}

	// check expired time
	if time.Now().After(user.OtpExpiredAt.Time) {
		return fiber.NewError(fiber.StatusBadRequest, "OTP is expired")
	}

	// clear OTP and expired time
	// user.Otp = ""
	// user.OtpExpiredAt = time.Time{}
	// isActive := true
	// user.IsActive = &isActive
	// n.UserRepository.Update(user)

	return nil
}

func (n *AuthServiceImpl) ResendOtp(email string) *fiber.Error {
	err := n.validate.Var(email, "required,email")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email format")
	}

	// Check if user exists in the database
	user, err := n.UserRepository.FindOneBy("email = ? AND otp IS NOT NULL", email)
	if err != nil || user == nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	//generate OTP
	otp, _ := utils.GenerateNumberOTP(6)

	user.Otp = sql.NullString{String: otp, Valid: true}
	user.OtpExpiredAt = sql.NullTime{Time: time.Now().Add(5 * time.Minute), Valid: true}
	user.OtpCounter++

	if user.OtpCounter >= 5 {
		if time.Now().Before(user.OtpExpiredAt.Time.Add(30 * time.Minute)) {
			return fiber.NewError(fiber.StatusBadRequest, "You have reached the maximum number of OTP requests. Please try again later.")
		}

		// reset OTP counter
		user.OtpCounter = 1
	}

	//send verification email
	serr := helpers.SendEmail(email, "Email verification", fmt.Sprintf("Your OTP: <br /> <h2>%s</h2>", otp))
	if serr != nil {
		panic(serr)
	}

	n.UserRepository.Update(user)
	return nil
}

func (n *AuthServiceImpl) ChangePassword(data request.ChangePasswordRequest, user *model.User) *fiber.Error {
	//validate input data
	if err := n.validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	if !user.CompareHashAndPassword(data.NewPassword) {
		return fiber.NewError(fiber.StatusBadRequest, "Wrong password")
	}

	newHash, gerr := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	helpers.ErrorPanic(gerr)
	user.Password = string(newHash)
	if uerr := n.UserRepository.Update(user); uerr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Change password unsuccessfully")
	}

	return nil
}
