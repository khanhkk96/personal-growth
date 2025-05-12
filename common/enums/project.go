package enums

import "github.com/go-playground/validator/v10"

type ProjectType string

const (
	WEB         ProjectType = "web"
	MOBILE_APP  ProjectType = "mobile_app"
	DESKTOP_APP ProjectType = "desktop_app"
	LIBRARY     ProjectType = "library"
)

// func IsValidProjectType(projectType ProjectType) bool {
// 	switch projectType {
// 	case WEB, MOBILE_APP, DESKTOP_APP, LIBRARY:
// 		return true

// 	default:
// 		return false
// 	}
// }

func isValidProjectType(fl validator.FieldLevel) bool {
	projectType := fl.Field().String()
	validTypes := map[string]bool{
		"web":         true,
		"mobile_app":  true,
		"desktop_app": true,
		"library":     true,
	}
	return validTypes[projectType]
}
