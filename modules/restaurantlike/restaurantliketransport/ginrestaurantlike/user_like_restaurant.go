package ginrestaurantlike

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/modules/restaurantlike/restaurantlikebiz"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"Delivery_Food/modules/restaurantlike/restaurantlikestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
