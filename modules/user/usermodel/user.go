package usermodel

import (
	"Delivery_Food/common"
	"Delivery_Food/component/tokenprovider"
	"errors"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"-" gorm:"column:password"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"-" gorm:"column:role"`
	Salt            string        `json:"-" gorm:"column:salt"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUserId() int {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (data *User) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"password" gorm:"column:password"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Role            string        `json:"-" gorm:"column:role"`
	Salt            string        `json:"-" gorm:"column:salt"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (data *UserCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email" form:"email"`
	Password string `json:"password" form:"password" gorm:"column:password form:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(accessToken, refreshToken *tokenprovider.Token) *Account {
	return &Account{accessToken, refreshToken}
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(errors.New(
		"usrname or password invalid"), "username or password invalid", "ErrUsernameOrPasswordInvalid")

	ErrEmailExisted = common.NewCustomError(errors.New("email has been used"),
		"email has been used",
		"ErrEmailExisted")
)
