package restaurantlikemodel

import (
	"Delivery_Food/common"
	"fmt"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at,omitempty" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string { return "restaurant_likes" }

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("can not like restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("can not unlike restaurant"),
		fmt.Sprintf("ErrCannotUnlikeRestaurant"),
	)
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

var (
	ErrUserLikedRestaurant = common.NewCustomError(nil, "the user liked this restaurant", "ErrUserLikedRestaurant")

	ErrUserDidNotLikeRestaurant = common.NewCustomError(nil, "the user did not like this restaurant", "ErrUserDidNotLikeRestaurant")
)
