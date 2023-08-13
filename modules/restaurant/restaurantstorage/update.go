package restaurantstorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"context"
	"gorm.io/gorm"
)

func (s *SqlStore) UpdateData(ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id= ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *SqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id= ?",
		id).Update("like_count", gorm.Expr("like_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *SqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id= ?",
		id).Update("like_count", gorm.Expr("like_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
