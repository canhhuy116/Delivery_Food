package restaurantlikestorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"context"
)

func (s *SqlStore) Create(ctx context.Context,
	data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
