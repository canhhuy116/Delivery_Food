package restaurantstorage

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"context"
)

func (s *SqlStore) ListDataByConditions(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	db := s.db

	var result []restaurantmodel.Restaurant

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(
		conditions).Where("status in (1)")

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
