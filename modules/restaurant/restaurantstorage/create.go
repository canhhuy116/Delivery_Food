package restaurantstorage

import (
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"golang.org/x/net/context"
)

func (s *SqlStore) Create(ctx context.Context,
	data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}