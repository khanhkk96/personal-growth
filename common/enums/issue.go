package enums

import "github.com/go-playground/validator/v10"

type IssuePriority = string

const (
	IP_LOW    IssuePriority = "low"
	IP_MEDIUM IssuePriority = "medium"
	IP_HIGH   IssuePriority = "high"
)

func isValidIssuePriority(fl validator.FieldLevel) bool {
	priorityLevel := fl.Field().String()
	validLevel := map[string]bool{
		"low":    true,
		"medium": true,
		"high":   true,
	}
	return validLevel[priorityLevel]
}

type IssueStatus string

const (
	IS_PENDING     IssueStatus = "pending"
	IS_PROCCESSING IssueStatus = "processing"
	IS_RESOLVED    IssueStatus = "resolved"
)

func isValidIssueStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	validStatus := map[string]bool{
		"pending":    true,
		"processing": true,
		"resolved":   true,
	}
	return validStatus[status]
}
