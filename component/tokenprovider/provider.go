package tokenprovider

import (
	"Delivery_Food/common"
	"errors"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)
	ErrEncoding = common.NewCustomError(
		errors.New("token encoding error"),
		"token encoding error",
		"ErrEncoding",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("token invalid"),
		"token invalid",
		"ErrInvalidToken",
	)
)

type Token struct {
	Token string `json:"token"`
	Created time.Time `json:"created"`
	Expiry int `json:"expiry"`
}

type TokenPayload struct {
	UserId int `json:"user_id"`
	Role string `json:"role"`
}