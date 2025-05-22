package responses

import (
	"personal-growth/common/enums"
	"time"
)

type IssueResponse struct {
	BaseResponse
	// Id          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Images      string              `json:"image"`
	ProjectId   string              `json:"project_id"`
	Project     ProjectResponse     `json:"project"`
	Status      enums.IssueStatus   `json:"status"`
	Priority    enums.IssuePriority `json:"priority"`
	IssuedAt    time.Time           `json:"issued_at"`
	NeedToSolve int                 `json:"need_to_solve"`
	References  string              `json:"references"`
}

type IssuePageResponse = BasePaginatedResponse[IssueResponse]
