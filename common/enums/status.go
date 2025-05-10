package enums

type ScheduleStatus string

const (
	PENDING    ScheduleStatus = "pending"
	PERFORMING ScheduleStatus = "performing"
	SUSPENDING ScheduleStatus = "suspending"
	DONE       ScheduleStatus = "done"
	CANCELED   ScheduleStatus = "canceled"
)

func IsValidScheduleStatus(status ScheduleStatus) bool {
	switch status {

	case PENDING, PERFORMING, SUSPENDING, DONE, CANCELED:
		return true

	default:
		return false
	}
}

type ProjectStatus string

const (
	PLANNING ProjectStatus = "planning"
	ONGOING  ProjectStatus = "ongoing"
	POSTPONE ProjectStatus = "postpone"
	FINISHED ProjectStatus = "finished"
)

func IsValidProjectStatus(status ProjectStatus) bool {
	switch status {
	case PLANNING, ONGOING, POSTPONE, FINISHED:
		return true
	default:
		return false
	}
}
