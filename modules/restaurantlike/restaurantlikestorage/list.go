package restaurantlikestorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"context"
)

func (s *SqlStore) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type tmp struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:like_count"`
	}

	var listLike []tmp

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).Select("restaurant_id, "+
		"COUNT(*) as like_count").
		Where("restaurant_id IN (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}
