package restaurantbiz

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"golang.org/x/net/context"
)

type ListRestaurantStore interface {
	ListDataByConditions(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type ListRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *ListRestaurantBiz {
	return &ListRestaurantBiz{store: store}
}

func (biz *ListRestaurantBiz) ListRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByConditions(ctx, conditions, filter, paging)

	return result, err
}
