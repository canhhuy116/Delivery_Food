package component

import (
	"Delivery_Food/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db, upProvider, secretKey}
}

func (ctx *appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
