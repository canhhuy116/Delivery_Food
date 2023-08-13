package ginrestaurantlike

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/modules/restaurantlike/restaurantlikebiz"
	"Delivery_Food/modules/restaurantlike/restaurantlikestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserUnlikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantlikebiz.NewUserUnlikeRestaurantBiz(store)

		if err := biz.UnlikeRestaurant(c.Request.Context(),
			requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
