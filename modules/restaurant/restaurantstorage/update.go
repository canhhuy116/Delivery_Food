package restaurantstorage

import (
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"golang.org/x/net/context"
)

func (s *SqlStore) UpdateData(ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id= ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}