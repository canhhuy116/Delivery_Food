package restaurantlikebiz

import (
	"Delivery_Food/component/asyncjob"
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

type DecreaseLikeCountRestaurantStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore DecreaseLikeCountRestaurantStore
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, decStore DecreaseLikeCountRestaurantStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store, decStore: decStore}
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

	go func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
