package ginuser

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/component/hasher"
	"Delivery_Food/component/tokenprovider/jwt"
	"Delivery_Food/modules/user/userbiz"
	"Delivery_Food/modules/user/usermodel"
	"Delivery_Food/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data usermodel.UserLogin

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDbConnection()
		tokenProvider := jwt.NewJwtProvider(appCtx.SecretKey())
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, md5, tokenProvider, 30*24*60*60)

		token, err := biz.Login(ctx.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
