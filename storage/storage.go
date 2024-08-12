package storage

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type CloudinaryClient struct {
	Logger logx.Logger
	Client *cloudinary.Cloudinary
	Folder string
}

func NewCloudinaryClient(CloudName, APIKey, APISecret, folder string) *CloudinaryClient {
	cld, err := cloudinary.NewFromParams(CloudName, APIKey, APISecret)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	return &CloudinaryClient{
		Logger: logx.WithContext(ctx),
		Client: cld,
		Folder: folder,
	}
}

func (c *CloudinaryClient) UploadImage(ctx context.Context, file *os.File) (string, error) {
	publicID := uuid.New().String()
	f, err := c.Client.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:  c.Folder,
		PublicID: publicID,
	})
	if err != nil {
		c.Logger.Errorf("failed to upload image: %v", err)
		return "", err
	}

	return f.SecureURL, nil
}