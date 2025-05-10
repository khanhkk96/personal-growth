package services

import (
	"database/sql"
	"fmt"
	"log"
	"personal-growth/common/constants"
	"personal-growth/configs"
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/helpers"
	"personal-growth/models"
	"personal-growth/repositories"
	"personal-growth/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repository repositories.UserRepository
	validate   *validator.Validate
}

func NewAuthServiceImpl(repository repositories.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		repository: repository,
		validate:   validate,
	}
}

func (n *AuthServiceImpl) Login(data requests.LoginRequest) (*responses.LoginResponse, *fiber.Error) {
	// Validate username and password
	if err := n.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	// Check if user exists in the database
	user, err := n.repository.FindOneBy("username = ? AND is_active = ?", data.Username, true)
	if err != nil || user == nil {
		log.Printf("User %s not found", data.Username)
		return nil, fiber.NewError(fiber.StatusNotFound, "Either username or password is incorrect")
	}

	// Check if password is correct
	if !user.CompareHashAndPassword(data.Password) {
		log.Println("Password is incorrect - username:", data.Username)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Either username or password is incorrect")
	}

	// generate access token
	config, _ := configs.LoadConfig(".")
	token, rf, err := utils.GenerateTokens(config.TokenExpiredIn, user.Id, config.TokenSecret, config.RefreshTokenSecret)
	helpers.ErrorPanic(err)

	return &responses.LoginResponse{
		AccessToken:  token,
		RefreshToken: rf,
	}, nil
}

func (n *AuthServiceImpl) RefreshAccessToken(refreshToken string) (string, *fiber.Error) {
	config, _ := configs.LoadConfig(".")
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

func (n *AuthServiceImpl) Register(data requests.RegisterRequest) (*models.User, *fiber.Error) {
	// Validate input data
	if err := n.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	// Check if user exists in the database
	user, _ := n.repository.FindOneBy("Username = ? OR Email = ? OR Phone = ?", data.Username, data.Email, data.Phone)
	if user != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "User exists")
	}

	//save user data
	user = &models.User{}
	copier.Copy(user, data)

	//generate OTP
	otp, _ := utils.GenerateNumberOTP(6)

	user.Otp = sql.NullString{String: otp, Valid: true}
	user.OtpExpiredAt = sql.NullTime{Time: time.Now().Add(5 * time.Minute), Valid: true}
	user.OtpCounter++

	cerr := n.repository.Create(user)
	if cerr != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Register your account unsuccessfully")
	}

	config, _ := configs.LoadConfig(".")
	content := helpers.RegistrationEmailData{
		Name:     user.FullName,
		AppName:  constants.APP_NAME,
		LoginURL: fmt.Sprintf("%s/login", config.ClientAddress),
		Otp:      otp,
	}

	message, err := helpers.RenderEmailTemplate("templates/welcome_email.html", content)
	if err != nil {
		log.Fatal(err)
	}

	//send verification email
	serr := helpers.SendEmail(data.Email, "Verify your account", message)
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
	user, err := n.repository.FindByEmail(email)
	if err != nil || user == nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if user.OtpCounter >= 5 {
		if time.Now().Before(user.OtpExpiredAt.Time.Add(3 * time.Minute)) {
			return fiber.NewError(fiber.StatusBadRequest, "You have reached the maximum number of OTP requests. Please try again later.")
		}

		// reset OTP counter
		user.OtpCounter = 0
	}

	//generate OTP
	otp, _ := utils.GenerateNumberOTP(6)
	content := helpers.RegistrationEmailData{
		AppName: constants.APP_NAME,
		Otp:     otp,
	}

	message, err := helpers.RenderEmailTemplate("templates/forgot_password_email.html", content)
	if err != nil {
		log.Fatal(err)
	}
	//send verification email
	serr := helpers.SendEmail(email, "Forgot password", message)
	if serr != nil {
		panic(serr)
	}

	user.Otp = sql.NullString{String: otp, Valid: true}
	user.OtpExpiredAt = sql.NullTime{Time: time.Now().Add(5 * time.Minute), Valid: true}
	user.OtpCounter++

	n.repository.Update(user)
	return nil
}

func (n *AuthServiceImpl) VerifyAccount(data requests.VerifyOTPRequest) *fiber.Error {
	// Check if user exists in the database
	err := n.VerifyOtp(data)
	if err != nil {
		return err
	}

	user, _ := n.repository.FindByEmail(data.Email)

	// clear OTP and expired time
	user.Otp = sql.NullString{Valid: false}
	user.OtpExpiredAt = sql.NullTime{Valid: false}
	user.OtpCounter = 0
	user.IsActive = true //activate account
	n.repository.Update(user)

	return nil
}

func (n *AuthServiceImpl) VerifyOtp(data requests.VerifyOTPRequest) *fiber.Error {
	// Validate input data
	if err := n.validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	// Check if user exists in the database
	user, err := n.repository.FindByEmail(data.Email)
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

	return nil
}

func (n *AuthServiceImpl) ResendOtp(email string) *fiber.Error {
	err := n.validate.Var(email, "required,email")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email format")
	}

	// Check if user exists in the database
	user, err := n.repository.FindOneBy("email = ? AND otp IS NOT NULL", email)
	if err != nil || user == nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	//generate OTP
	otp, _ := utils.GenerateNumberOTP(6)

	if user.OtpCounter >= 5 {
		if time.Now().Before(user.OtpExpiredAt.Time.Add(3 * time.Minute)) {
			return fiber.NewError(fiber.StatusBadRequest, "You have reached the maximum number of OTP requests. Please try again later.")
		}

		// reset OTP counter
		user.OtpCounter = 0
	}

	user.Otp = sql.NullString{String: otp, Valid: true}
	user.OtpExpiredAt = sql.NullTime{Time: time.Now().Add(5 * time.Minute), Valid: true}
	user.OtpCounter++

	//send verification email
	serr := helpers.SendEmail(email, "Email verification", fmt.Sprintf("<p>Your OTP: </p><h2>%s</h2>", otp))
	if serr != nil {
		panic(serr)
	}

	n.repository.Update(user)
	return nil
}

func (n *AuthServiceImpl) ChangePassword(data requests.ChangePasswordRequest, user *models.User) *fiber.Error {
	//validate input data
	if err := n.validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	if !user.CompareHashAndPassword(data.OldPassword) {
		return fiber.NewError(fiber.StatusBadRequest, "Incorrect password")
	}

	newHash, gerr := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	helpers.ErrorPanic(gerr)
	user.Password = string(newHash)
	if uerr := n.repository.Update(user); uerr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Change password unsuccessfully")
	}

	return nil
}

func (n *AuthServiceImpl) SetNewPassword(data requests.SetNewPasswordRequest) *fiber.Error {
	//validate input data
	if err := n.validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid data")
	}

	verr := n.VerifyOtp(requests.VerifyOTPRequest{Otp: data.Otp, Email: data.Email})
	if verr != nil {
		return verr
	}

	user, _ := n.repository.FindByEmail(data.Email)

	// clear OTP and expired time
	user.Otp = sql.NullString{Valid: false}
	user.OtpExpiredAt = sql.NullTime{Valid: false}
	user.OtpCounter = 0

	newHash, gerr := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	helpers.ErrorPanic(gerr)
	user.Password = string(newHash)
	if uerr := n.repository.Update(user); uerr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Change password unsuccessfully")
	}

	return nil
}

func (n *AuthServiceImpl) UploadAvatar(file string, user *models.User) *fiber.Error {
	// Validate input data
	if file == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid avatar")
	}

	user.Avatar = sql.NullString{String: file, Valid: true}
	if uerr := n.repository.Update(user); uerr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Upload avatar unsuccessfully")
	}

	return nil
}
