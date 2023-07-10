package main

import (
	"Delivery_Food/component"
	"Delivery_Food/middleware"
	"Delivery_Food/modules/restaurant/restauranttransport/ginrestaurant"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

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
