package restaurantbiz

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"context"
	"log"
)

type ListRestaurantStore interface {
	ListDataByConditions(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type ListRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantBiz(store ListRestaurantStore,
	likeStore LikeStore) *ListRestaurantBiz {
	return &ListRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *ListRestaurantBiz) ListRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByConditions(ctx, conditions, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].ID
	}

	likeMap, err := biz.likeStore.GetRestaurantLike(ctx, ids)

	if err != nil {
		log.Println("Error at GetRestaurantLike: ", err)
	}

	if v := likeMap; v != nil {
		for i, item := range result {
			result[i].LikeCount = likeMap[item.ID]
		}
	}

	return result, nil
}
