package helpers

import (
	"fmt"
	"personal-growth/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/thoas/go-funk"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func CustomErrorPanic(err error, msg string) {
	if err != nil {
		panic(utils.Ternary(funk.IsEmpty(msg), err.Error(), msg))
	}
}

type ExceptionError struct {
	Code    int
	Message string
}

func (e *ExceptionError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func PrintErrorMessage(err error) string {
	if err != nil {
		errors := make([]string, 0)

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errors = append(errors, fmt.Sprintf("%s does not meet the '%s' requirement", e.Field(), strings.ToUpper(e.Tag())))
			}
		} else {
			errors = append(errors, fmt.Sprintln("Validation error:", err))
		}

		return strings.Join(errors, ", ")
	}

	return ""
}
