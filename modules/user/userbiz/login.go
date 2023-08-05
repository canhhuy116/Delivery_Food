package userbiz

import (
	"Delivery_Food/common"
	"Delivery_Food/component"
	"Delivery_Food/component/tokenprovider"
	"Delivery_Food/modules/user/usermodel"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

//type TokenConfig interface {
//	GetAtExp() int
//	GetRtExp() int
//}

type loginBiz struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
	//tokenCfg      TokenConfig
}

func NewLoginBiz(storeUser LoginStorage, hasher Hasher,
	tokenProvider tokenprovider.Provider, expiry int) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		hasher:        hasher,
		tokenProvider: tokenProvider,
		//tokenCfg:      tokenCfg,
		expiry: expiry,
	}
}

// 1. Find user by email
// 2. Hash password from request and compare with password from db
// 3. Generate access token and refresh token
// 4. Save refresh token to db
// 5. Return access token and refresh token

func (biz *loginBiz) Login(ctx context.Context,
	data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.ID,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := biz.tokenProvider.Generate(payload, biz.tokenCfg.GetRtExp())
	//
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}
	return accessToken, nil
}
