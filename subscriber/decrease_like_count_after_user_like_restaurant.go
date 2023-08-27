package subscriber

import (
	"Delivery_Food/component"
	"Delivery_Food/modules/restaurant/restaurantstorage"
	"Delivery_Food/pubsub"
	"context"
)

func RunDecreaseLikeCountAfterUserLikeRestaurant(appCtx component.
	AppContext) ConsumerJob {
	return ConsumerJob{
		Title: "Decrease like count after user unlike restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
			data := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, data.GetRestaurantId())
		},
	}
}
