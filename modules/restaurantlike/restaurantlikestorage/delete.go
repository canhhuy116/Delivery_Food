package restaurantlikestorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"context"
)

func (s *SqlStore) Delete(ctx context.Context,
	userId, restaurantId int) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
