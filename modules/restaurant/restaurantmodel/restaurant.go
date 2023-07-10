package restaurantmodel

import (
	"Delivery_Food/common"
	"errors"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"id" gorm:"column:id"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"address" gorm:"addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"address" gorm:"addr"`
}

func (*RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string `json:"name" gorm:"column:name"`
	Addr            *string `json:"address" gorm:"addr"`
}

func (*RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}
