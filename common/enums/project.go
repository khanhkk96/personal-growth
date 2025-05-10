package enums

type ProjectType string

const (
	WEB         ProjectType = "web"
	MOBILE_APP  ProjectType = "mobile_app"
	DESKTOP_APP ProjectType = "desktop_app"
	LIBRARY     ProjectType = "library"
)

func IsValidProjectType(projectType ProjectType) bool {
	switch projectType {
	case WEB, MOBILE_APP, DESKTOP_APP, LIBRARY:
		return true

	default:
		return false
	}
}
