package uploadstorage

import (
	"Delivery_Food/common"
	"context"
)

func (s *SqlStore) CreateImage(ctx context.Context, data *common.Image) error {
	db := s.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
