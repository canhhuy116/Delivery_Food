package ginrestaurant

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/modules/restaurant/restaurantbiz"
	"Delivery_Food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantbiz.NewFindRestaurantBiz(store)
		result, err := biz.FindRestaurant(c.Request.Context(),
			map[string]interface{}{"id": uid.GetLocalID()})

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
