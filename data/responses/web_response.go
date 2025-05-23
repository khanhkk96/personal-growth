package responses

import (
	"math"
	"time"

	"github.com/google/uuid"
)

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

type BaseResponse struct {
	Id          string        `json:"id"`
	CreatedById uuid.UUID     `json:"created_by_id"`
	CreatedBy   *UserResponse `json:"created_by" type:"UserResponse"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type PaginationMetaData[T any] struct {
	page       int
	limit      int
	totalItems int
	data       []T
}

func NewPaginationMetaData[T any](page int, limit int, totalItems int, data []T) PaginationMetaData[T] {
	return PaginationMetaData[T]{
		page:       page,
		limit:      limit,
		totalItems: totalItems,
		data:       data,
	}
}

func NewPaginatedResponse[T any](meta PaginationMetaData[T]) BasePaginatedResponse[T] {
	// Calculate total pages using math.Ceil for proper rounding
	totalPages := int(math.Ceil(float64(meta.totalItems) / float64(meta.limit)))

	return BasePaginatedResponse[T]{
		Page:            meta.page,
		Limit:           meta.limit,
		TotalItems:      meta.totalItems,
		PageCount:       totalPages,
		HasPreviousPage: meta.page > 1,
		HasNextPage:     meta.page < totalPages,
		Items:           meta.data,
	}
}
