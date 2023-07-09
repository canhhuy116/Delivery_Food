package restaurantbiz

import (
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"golang.org/x/net/context"
)

type FindRestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type FindRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFindRestaurantBiz(store FindRestaurantStore) *FindRestaurantBiz {
	return &FindRestaurantBiz{store: store}
}

func (biz *FindRestaurantBiz) FindRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindByConditions(ctx, conditions)

	return result, err
}
