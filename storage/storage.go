package storage

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
)

type StorageClient struct {
	logx.Logger
	Client    *minio.Client
	ParentDir string
}

func NewStorageClient(endpoint, accessKey, secretKey string) *StorageClient {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	return &StorageClient{
		Logger: logx.WithContext(ctx),
		Client: client,
	}
}

func (sc *StorageClient) PushFile(ctx context.Context, name string, data *bytes.Buffer) (string, error) {
	baseBucket := "tenant-6"

	_, err := sc.Client.PutObject(ctx, baseBucket, "temp/"+name, data, int64(data.Len()), minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	return baseBucket + "/temp/" + name, nil
}
