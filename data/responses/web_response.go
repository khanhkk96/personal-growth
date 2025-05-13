package responses

import (
	"math"

	"github.com/jinzhu/copier"
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

type PaginationMetaData struct {
	page       int
	limit      int
	totalItems int
	data       []interface{}
}

func NewPaginationMetaData(page int, limit int, totalItems int, data []interface{}) PaginationMetaData {
	return PaginationMetaData{
		page:       page,
		limit:      limit,
		totalItems: totalItems,
		data:       data,
	}
}

func NewPaginatedResponse[T any](meta PaginationMetaData) BasePaginatedResponse[T] {
	// Calculate total pages using math.Ceil for proper rounding
	totalPages := int(math.Ceil(float64(meta.totalItems) / float64(meta.limit)))

	var items []T
	for _, value := range meta.data {
		var item T
		// Handle potential errors from copier.Copy
		if err := copier.Copy(&item, value); err != nil {
			continue // Skip the item if copying fails
		}
		items = append(items, item)
	}

	return BasePaginatedResponse[T]{
		Page:            meta.page,
		Limit:           meta.limit,
		TotalItems:      meta.totalItems,
		PageCount:       totalPages,
		HasPreviousPage: meta.page > 1,
		HasNextPage:     meta.page < totalPages,
		Items:           items,
	}
}
