package enums

import "github.com/go-playground/validator/v10"

type ScheduleStatus string

const (
	SS_PENDING    ScheduleStatus = "pending"
	SS_PERFORMING ScheduleStatus = "performing"
	SS_SUSPENDING ScheduleStatus = "suspending"
	SS_DONE       ScheduleStatus = "done"
	SS_CANCELED   ScheduleStatus = "canceled"
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
	PS_PLANNING ProjectStatus = "planning"
	PS_ONGOING  ProjectStatus = "ongoing"
	PS_POSTPONE ProjectStatus = "postpone"
	PS_FINISHED ProjectStatus = "finished"
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
