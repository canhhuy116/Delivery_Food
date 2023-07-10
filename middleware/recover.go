package middleware

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"github.com/gin-gonic/gin"
)

func Recover(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				// Check if the error is an instance of AppError
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err) // Catch panic of gin
					return
				}

				// If the error is not an AppError, create a custom error and return it
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err) // Catch panic of gin
				return
			}
		}()

		// Call the next handler or middleware
		c.Next()
	}
}
