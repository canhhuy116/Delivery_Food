package restaurantlikebiz

import (
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"context"
)

type UserLikeRestaurantStore interface {
	FindByConditions(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantlikemodel.Like, error)
	Create(ctx context.Context,
		data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	liked, _ := biz.store.FindByConditions(ctx, map[string]interface{}{
		"user_id":       data.UserId,
		"restaurant_id": data.RestaurantId,
	})

	if liked != nil {
		return restaurantlikemodel.ErrUserLikedRestaurant
	}

	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
