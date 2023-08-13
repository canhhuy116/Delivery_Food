package restaurantlikebiz

import (
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"context"
)

type UserUnlikeRestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantlikemodel.Like, error)
	Delete(ctx context.Context,
		userId, restaurantId int) error
}

type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId, restaurantId int,
) error {
	liked, _ := biz.store.FindByConditions(ctx, map[string]interface{}{
		"user_id":       userId,
		"restaurant_id": restaurantId,
	})

	if liked == nil {
		return restaurantlikemodel.ErrUserDidNotLikeRestaurant
	}

	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	return nil
}
