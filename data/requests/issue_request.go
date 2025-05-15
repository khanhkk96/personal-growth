package requests

import (
	"personal-growth/common/enums"
	"time"
)

type CreateOrUpdateIssueRequest struct {
	Name        string              `json:"name" validate:"required,min=6,max=256"`
	Description string              `json:"description" validate:"max=500,omitempty"`
	ProjectId   string              `json:"project_id" validate:"omitempty"`
	Status      enums.IssueStatus   `json:"status" validate:"issue_status_enum"`
	Priority    enums.IssuePriority `json:"priority" validate:"issue_priority_enum"`
	IssuedAt    *time.Time          `json:"issued_at"`
	NeedToSolve int                 `json:"need_to_solve"`
	References  string              `json:"references" validate:"max:500,omitempty"`
}

type IssueFilters struct {
	BaseRequest
	Status   *enums.IssueStatus   `json:"status" validate:"issue_status_enum"`
	Priority *enums.IssuePriority `json:"priority" validate:"issue_priority_enum"`
}
