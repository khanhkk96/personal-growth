package enums

type UserType string

const (
	ADMIN UserType = "admin"
	USER  UserType = "user"
)

func IsValidUserType(userType UserType) bool {
	switch userType {
	case ADMIN, USER:
		return true
	default:
		return false
	}
}
