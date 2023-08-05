package userstorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/user/usermodel"
	"context"
)

func (s *SqlStore) FindUser(ctx context.Context,
	conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var data usermodel.User

	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
