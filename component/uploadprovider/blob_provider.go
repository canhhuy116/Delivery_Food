package uploadprovider

import (
	"Delivery_Food/common"
	"context"
	"fmt"
	_ "log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type AzureBlobProvider struct {
	accountName         string
	accountKey          string
	containerName       string
	containerURL        azblob.ContainerURL
	containerURLWithSAS string
	domain              string
}

func NewAzureBlobProvider(accountName, accountKey, containerName, domain string) (*AzureBlobProvider, error) {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, err := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
	if err != nil {
		return nil, err
	}

	containerURL := azblob.NewContainerURL(*URL, pipeline)
	containerURLWithSAS := containerURL.URL()
	if strings.Contains(containerURLWithSAS.RawQuery, "?") {
		containerURLWithSAS.RawQuery += "&"
	} else {
		containerURLWithSAS.RawQuery += "?"
	}
	containerURLWithSAS.RawQuery += "sv=2019-12-12&ss=b&srt=sco&sp=rwdlacx&se=2025-01-01T00:00:00Z&st=2021-01-01T00:00:00Z"

	return &AzureBlobProvider{
		accountName:         accountName,
		accountKey:          accountKey,
		containerName:       containerName,
		containerURL:        containerURL,
		containerURLWithSAS: containerURLWithSAS.String(),
		domain:              domain,
	}, nil
}

func (provider *AzureBlobProvider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	blobURL := provider.containerURL.NewBlockBlobURL(dst)

	_, err := azblob.UploadBufferToBlockBlob(ctx, data, blobURL, azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: http.DetectContentType(data),
		},
	})
	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url: fmt.Sprintf("%s/%s/%s", provider.domain,
			provider.containerName, dst),
		CloudName: "azure",
	}

	return img, nil
}
