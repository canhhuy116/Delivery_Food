package subscriber

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/modules/restaurant/restaurantstorage"
	"Delivery_Food/pubsub"
	"context"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext,
	ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
	store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())

	go func() {
		defer common.AppRecover()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-c:
				if msg == nil {
					return
				}

				data := msg.Data().(HasRestaurantId)
				if err := store.IncreaseLikeCount(ctx, data.GetRestaurantId()); err != nil {
					panic(err)
				}
			}
		}
	}()
}

//func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.
//	AppContext) func(ctx context.Context, message pubsub.Message) error {
//	store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
//
//	return func(ctx context.Context,message pubsub.Message) error {
//		data := message.Data().(HasRestaurantId)
//		return store.IncreaseLikeCount(ctx,
//			data.GetRestaurantId())
//	}
//}

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.
	AppContext) ConsumerJob {
	return ConsumerJob{
		Title: "Increase like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
			data := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, data.GetRestaurantId())
		},
	}
}
