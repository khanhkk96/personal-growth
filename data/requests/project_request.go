package requests

import (
	"personal-growth/common/enums"
	"time"
)

type CreateOrUpdateProjectRequest struct {
	Name        string              `json:"name" validate:"required,min=1,max=256"`
	Type        enums.ProjectType   `json:"type" validate:"required,project_type_enum"`
	Summary     string              `json:"summary" validate:"required,max=500"`
	Description string              `json:"description" validate:"max=500"`
	Stack       string              `json:"stack" validate:"required,min=1,max=256"`
	StartAt     *time.Time          `json:"start_at"`
	EndAt       *time.Time          `json:"end_at"`
	Status      enums.ProjectStatus `json:"status" validate:"required,project_status_enum"`
}

type ProjectFilters struct {
	BaseRequest
	Status *enums.ProjectStatus `json:"status" validate:"project_status_enum"`
	Type   *enums.ProjectType   `json:"type" validate:"project_type_enum"`
}
