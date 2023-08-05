package userstorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/user/usermodel"
	"context"
)

func (s *SqlStore) CreateUser(ctx context.Context,
	data *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
