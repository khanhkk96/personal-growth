package enums

import "github.com/go-playground/validator/v10"

type ScheduleStatus string

const (
	PENDING    ScheduleStatus = "pending"
	PERFORMING ScheduleStatus = "performing"
	SUSPENDING ScheduleStatus = "suspending"
	DONE       ScheduleStatus = "done"
	CANCELED   ScheduleStatus = "canceled"
)

// func IsValidScheduleStatus(status ScheduleStatus) bool {
// 	switch status {

// 	case PENDING, PERFORMING, SUSPENDING, DONE, CANCELED:
// 		return true

// 	default:
// 		return false
// 	}
// }

func isValidScheduleStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	validStatus := map[string]bool{
		"pending":    true,
		"performing": true,
		"suspending": true,
		"done":       true,
		"canceled":   true,
	}
	return validStatus[status]
}

type ProjectStatus string

const (
	PLANNING ProjectStatus = "planning"
	ONGOING  ProjectStatus = "ongoing"
	POSTPONE ProjectStatus = "postpone"
	FINISHED ProjectStatus = "finished"
)

// func IsValidProjectStatus(status ProjectStatus) bool {
// 	switch status {
// 	case PLANNING, ONGOING, POSTPONE, FINISHED:
// 		return true
// 	default:
// 		return false
// 	}
// }

func isValidProjectStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	validStatus := map[string]bool{
		"planning": true,
		"ongoing":  true,
		"postpone": true,
		"finished": true,
	}
	return validStatus[status]
}
