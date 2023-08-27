package restaurantlikebiz

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"Delivery_Food/pubsub"
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
	store UserUnlikeRestaurantStore
	//decStore DecreaseLikeCountRestaurantStore
	pubsub pubsub.PubSub
}

func NewUserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
	//decStore DecreaseLikeCountRestaurantStore,
	pubsub pubsub.PubSub,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
		//decStore: decStore,
		pubsub: pubsub,
	}
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

	biz.pubsub.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(liked))

	return nil
}
