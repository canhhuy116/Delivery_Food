package restaurantbiz

import (
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"errors"
	"golang.org/x/net/context"
)

type DeleteRestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)

	SoftDeleteData(ctx context.Context, id int) error
}

type DeleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *DeleteRestaurantBiz {
	return &DeleteRestaurantBiz{store: store}
}

func (biz *DeleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindByConditions(ctx,
		map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}

	return nil
}
