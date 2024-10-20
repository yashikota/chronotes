package minio

import (
	"context"
	"log/slog"

	"github.com/minio/minio-go/v7"
)

var MinioCtx context.Context
var bucketName = "chronotes"

func CreateBucket() {
	MinioCtx = context.Background()

	location := "ap-northeast-1"

	err := MinioClient.MakeBucket(MinioCtx, bucketName, minio.MakeBucketOptions{
		Region: location,
	})
	if err != nil {
		exists, errBucketExits := MinioClient.BucketExists(MinioCtx, bucketName)
		if errBucketExits == nil && exists {
			slog.Warn("we already own " + bucketName)
		} else {
			slog.Error("err: " + err.Error())
		}
	} else {
		slog.Info("successfully created bucket: " + bucketName)
	}
}
