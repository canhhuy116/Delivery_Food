package restaurantbiz

import (
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"errors"
	"golang.org/x/net/context"
)

type UpdateRestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)

	UpdateData(ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate) error
}

type UpdateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *UpdateRestaurantBiz {
	return &UpdateRestaurantBiz{store: store}
}

func (biz *UpdateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	oldData, err := biz.store.FindByConditions(ctx,
		map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
