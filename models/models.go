package models

import (
	"math"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/pkg"
	"gorm.io/gorm"
)

func paginate(value interface{}, pagination *pkg.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows

	limit := pagination.Limit
	if limit <= 0 {
		limit = 1
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
