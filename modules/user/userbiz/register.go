package userbiz

import (
	"Delivery_Food/common"
	"Delivery_Food/modules/user/usermodel"
	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type RegisterBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBiz(store RegisterStorage, hasher Hasher) *RegisterBiz {
	return &RegisterBiz{
		registerStorage: store,
		hasher:          hasher,
	}
}

func (biz *RegisterBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.registerStorage.FindUser(ctx,
		map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
