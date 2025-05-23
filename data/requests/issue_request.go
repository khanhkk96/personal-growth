package requests

import (
	"personal-growth/common/enums"
	"time"
)

type CreateOrUpdateIssueRequest struct {
	Name        string              `form:"name" validate:"required,min=6,max=256"`
	Description string              `form:"description" validate:"max=500,omitempty"`
	ProjectId   string              `form:"project_id"`
	Status      enums.IssueStatus   `form:"status" validate:"required,issue_status_enum"`
	Priority    enums.IssuePriority `form:"priority" validate:"required,issue_priority_enum"`
	IssuedAt    *time.Time          `form:"issued_at"`
	NeedToSolve *int                `form:"need_to_solve"`
	References  string              `form:"references" validate:"max=500,omitempty"`
}

type IssueFilters struct {
	BaseRequest
	Status   *enums.IssueStatus   `query:"status" validate:"issue_status_enum"`
	Priority *enums.IssuePriority `query:"priority" validate:"issue_priority_enum"`
}
