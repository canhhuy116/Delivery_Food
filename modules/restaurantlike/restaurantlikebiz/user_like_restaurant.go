package restaurantlikebiz

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"Delivery_Food/pubsub"
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
	store UserLikeRestaurantStore
	//incStore IncreaseLikeCountRestaurantStore
	pubsub pubsub.PubSub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	//incStore IncreaseLikeCountRestaurantStore,
	pubsub pubsub.PubSub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		//incStore: incStore,
		pubsub: pubsub,
	}
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

	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant,
		pubsub.NewMessage(data))

	//go func() {
	//	defer common.AppRecover()
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//	})
	//
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	return nil
}
