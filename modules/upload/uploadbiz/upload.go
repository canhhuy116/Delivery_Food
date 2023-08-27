package uploadbiz

import (
	"Delivery_Food/common"
	"Delivery_Food/component/uploadprovider"
	"Delivery_Food/modules/upload/uploadmodel"
	"bytes"
	"context"
	"fmt"
	"image"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type UploadBiz struct {
	provider uploadprovider.UploadProvider
	store    CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, store CreateImageStorage) *UploadBiz {
	return &UploadBiz{provider: provider, store: store}
}

func (biz *UploadBiz) Upload(ctx context.Context, folder string,
	fileName string, data []byte) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func getImageDimension(fileBytes *bytes.Buffer) (int, int, error) {
	img, _, err := image.DecodeConfig(fileBytes)

	if err != nil {
		log.Println("Error decode image config: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
