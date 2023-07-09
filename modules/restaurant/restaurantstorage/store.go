package restaurantstorage

import "gorm.io/gorm"

// SqlStore create struct store db
type SqlStore struct {
	db *gorm.DB
}

// NewSQLStore declare function NewSQLStore return sqlStore
func NewSQLStore(db *gorm.DB) *SqlStore {
	return &SqlStore{db}
}
