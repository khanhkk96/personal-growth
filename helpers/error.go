package helpers

import (
	"fmt"
	"personal-growth/utils"

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
