package ginupload

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/modules/upload/uploadbiz"
	"Delivery_Food/modules/upload/uploadstorage"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDbConnection()

		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := ctx.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
		}(file)

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), imgStore)
		img, err := biz.Upload(ctx.Request.Context(), folder,
			fileHeader.Filename, dataBytes)

		if err != nil {
			panic(err)
		}

		ctx.JSON(200, common.SimpleSuccessResponse(img))
	}
}
