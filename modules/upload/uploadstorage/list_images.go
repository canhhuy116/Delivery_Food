package uploadstorage

import (
	"Delivery_Food/common"
	"context"
)

func (s *SqlStore) ListImages(ctx context.Context, ids []int,
	moreKeys ...string,
) ([]common.Image, error) {
	db := s.db

	var images []common.Image
	if err := db.Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Find(&images).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return images, nil
}
