package utils

import (
	"math"
	"template-go-api/models"
)

// CalculatePagination calculates pagination metadata
func CalculatePagination(page, limit int, total int64) models.PaginationResponse {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return models.PaginationResponse{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: total,
		HasNext:      page < totalPages,
		HasPrev:      page > 1,
	}
}

// GetOffset calculates the offset for database queries
func GetOffset(page, limit int) int {
	return (page - 1) * limit
}
