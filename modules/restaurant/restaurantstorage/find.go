package restaurantstorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"context"
	"gorm.io/gorm"
)

func (s *SqlStore) FindByConditions(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	db := s.db

	var result restaurantmodel.Restaurant

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	err := db.First(&result).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
