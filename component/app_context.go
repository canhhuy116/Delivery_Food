package component

import (
	"Delivery_Food/component/uploadprovider"
	"Delivery_Food/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
	pb         pubsub.PubSub
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider,
	secretKey string, pb pubsub.PubSub) *appCtx {
	return &appCtx{db, upProvider, secretKey, pb}
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

func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.pb
}
