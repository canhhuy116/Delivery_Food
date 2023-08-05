package userstorage

import "gorm.io/gorm"

type SqlStore struct {
	db *gorm.DB
}

// NewSQLStore declare function NewSQLStore return sqlStore
func NewSQLStore(db *gorm.DB) *SqlStore {
	return &SqlStore{db}
}
