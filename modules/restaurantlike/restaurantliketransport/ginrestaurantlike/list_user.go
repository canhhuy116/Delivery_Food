package ginrestaurantlike

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/modules/restaurant/restaurantmodel"
	"Delivery_Food/modules/restaurantlike/restaurantlikebiz"
	"Delivery_Food/modules/restaurantlike/restaurantlikemodel"
	"Delivery_Food/modules/restaurantlike/restaurantlikestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDbConnection())
		biz := restaurantlikebiz.NewListUsersLikeRestaurantBiz(store)

		result, err := biz.ListUsersLikeRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(common.ErrCannotListEntity(restaurantmodel.EntityName, err))
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
