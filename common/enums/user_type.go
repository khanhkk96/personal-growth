package enums

import "github.com/go-playground/validator/v10"

type UserType string

const (
	ADMIN UserType = "admin"
	USER  UserType = "user"
)

// func IsValidUserType(userType UserType) bool {
// 	switch userType {
// 	case ADMIN, USER:
// 		return true
// 	default:
// 		return false
// 	}
// }

func isValidRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	validRoles := map[string]bool{
		"admin": true,
		"user":  true,
	}
	return validRoles[role]
}
