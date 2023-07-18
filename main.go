package main

import (
	"Delivery_Food/component"
	"Delivery_Food/component/uploadprovider"
	"Delivery_Food/middleware"
	"Delivery_Food/modules/restaurant/restauranttransport/ginrestaurant"
	"Delivery_Food/modules/upload/uploadtransport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBConnectionStr")

	//s3BucketName := os.Getenv("S3BucketName")
	//s3Region := os.Getenv("S3Region")
	//s3APIKey := os.Getenv("S3APIKey")
	//s3SecretKey := os.Getenv("S3SecretKey")
	//s3Domain := os.Getenv("S3Domain")
	//
	//s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey,
	//	s3SecretKey, s3Domain)

	accountName := os.Getenv("AzureAccountName")
	accountKey := os.Getenv("AzureAccountKey")
	containerName := os.Getenv("AzureContainerName")
	domain := os.Getenv("AzureDomain")

	azureProvider, err := uploadprovider.NewAzureBlobProvider(accountName, accountKey, containerName, domain)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//if err := runService(db, s3Provider); err != nil {
	//	log.Fatalln(err)
	//}
	if err := runService(db, azureProvider); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/upload", ginupload.Upload(appCtx))

	restaurant := r.Group("/restaurants")
	{
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurant.GET("/list", ginrestaurant.ListRestaurant(appCtx))
		restaurant.GET("/:id", ginrestaurant.FindRestaurant(appCtx))
		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run()
}
