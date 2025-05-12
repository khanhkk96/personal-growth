package enums

import "github.com/go-playground/validator/v10"

func RegisterCustomValidations(v *validator.Validate) {
	v.RegisterValidation("user_type_enum", isValidRole)
	v.RegisterValidation("schedule_status_enum", isValidScheduleStatus)
	v.RegisterValidation("project_status_enum", isValidProjectStatus)
	v.RegisterValidation("project_type_enum", isValidProjectType)
}
