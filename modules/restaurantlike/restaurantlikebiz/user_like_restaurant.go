package restaurantlikebiz

import (
	"Delivery_Food/component/asyncjob"
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

type IncreaseLikeCountRestaurantStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncreaseLikeCountRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncreaseLikeCountRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
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

	go func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
