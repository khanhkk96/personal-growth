package responses

import (
	"personal-growth/common/enums"
	"time"
)

type ProjectResponse struct {
	// Id          string              `json:"id"`
	BaseResponse
	Name        string              `json:"name"`
	Type        enums.ProjectType   `json:"type"`
	Summary     string              `json:"summary"`
	Stack       string              `json:"stack"`
	Description string              `json:"description"`
	StartAt     *time.Time          `json:"start_at"`
	EndAt       *time.Time          `json:"end_at"`
	Status      enums.ProjectStatus `json:"status"`
}

type ProjectPageResponse = BasePaginatedResponse[ProjectResponse]
