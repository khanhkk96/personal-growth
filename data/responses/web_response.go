package responses

import "math"

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type BasePaginatedResponse[T any] struct {
	Items           []T  `json:"items"`
	Page            int  `json:"page"`
	Limit           int  `json:"limit"`
	TotalItems      int  `json:"totalItems"`
	PageCount       int  `json:"pageCount"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	HasNextPage     bool `json:"hasNextPage"`
}

func NewPaginatedResponse[T any](page, limit, totalItems int, data []T) BasePaginatedResponse[T] {
	totalPages := int(math.Round(float64(totalItems / limit)))
	return BasePaginatedResponse[T]{
		Page:            page,
		Limit:           limit,
		TotalItems:      totalItems,
		PageCount:       totalPages,
		HasPreviousPage: page > 1,
		HasNextPage:     page < totalPages,
		Items:           data,
	}
}
