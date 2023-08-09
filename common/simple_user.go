package common

const EntityName = "User"

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (data *SimpleUser) Mask(isAdminOrOwner bool) {
	data.GenUID(DbUser)
}
