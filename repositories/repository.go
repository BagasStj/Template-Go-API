package repositories

import (
	"template-go-api/logger"

	"gorm.io/gorm"
)

type Repository struct {
	logger  logger.Logger
	dbRead  *gorm.DB
	dbWrite *gorm.DB
}

// NewRepository creates a new instance of Repository.
func NewRepository(logger logger.Logger, dbRead, dbWrite *gorm.DB) *Repository {
	return &Repository{
		logger:  logger,
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}

// GetReadDB returns the read database instance
func (r *Repository) GetReadDB() *gorm.DB {
	return r.dbRead
}

// GetWriteDB returns the write database instance
func (r *Repository) GetWriteDB() *gorm.DB {
	return r.dbWrite
}
