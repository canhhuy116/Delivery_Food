package restaurantlikestorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *SqlStore) FindByConditions(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*restaurantlikemodel.Like, error) {
	db := s.db

	var result restaurantlikemodel.Like

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	err := db.First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
