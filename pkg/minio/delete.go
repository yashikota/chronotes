package minio

import (
	"github.com/minio/minio-go/v7"
)

func DeleteObject(objectName string) error {
	err := MinioClient.RemoveObject(MinioCtx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
