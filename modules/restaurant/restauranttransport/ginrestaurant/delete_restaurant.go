package ginrestaurant

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/modules/restaurant/restaurantbiz"
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"Delivery_Food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(),
			int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
