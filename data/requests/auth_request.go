package requests

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type RegisterRequest struct {
	Username string `validate:"required,min=6,max=50" json:"username"`
	Password string `validate:"required,min=6,max=20" json:"password"`
	FullName string `validate:"required,min=6,max=100" json:"full_name"`
	Email    string `validate:"required,email" json:"email"`
	Phone    string `validate:"min=9,max=15" json:"phone"`
}

type ChangePasswordRequest struct {
	OldPassword string `validate:"required,min=6,max=20" json:"old_password"`
	NewPassword string `validate:"required,min=6,max=20" json:"new_password"`
}

type SetNewPasswordRequest struct {
	Email       string `validate:"required,email" json:"email"`
	NewPassword string `validate:"required,min=6,max=20" json:"new_password"`
	Otp         string `validate:"required,min=6,max=8" json:"otp"`
}

type ForgotPasswordRequest struct {
	Email string `validate:"required,email" json:"email"`
}

type ResendOTPRequest struct {
	Email string `validate:"required,email" json:"email"`
}

type VerifyOTPRequest struct {
	Email string `validate:"required,email" json:"email"`
	Otp   string `validate:"required,min=6,max=8" json:"otp"`
}
