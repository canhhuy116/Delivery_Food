package subscriber

import (
	"Delivery_Food/component"
	"context"
)

func Setup(appCtx component.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
	//DecreaseLikeCountAfterUserDislikeRestaurant(appCtx, context.Background())
}
