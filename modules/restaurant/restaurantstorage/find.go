package restaurantstorage

import (
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"golang.org/x/net/context"
)

func (s *SqlStore) FindByConditions(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	db := s.db

	var result *restaurantmodel.Restaurant

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
