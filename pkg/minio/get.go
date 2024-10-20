package minio

import (
	"time"
)

func GetObjectURL(objectName string) (string, error) {
	url, err := MinioClient.PresignedGetObject(MinioCtx, bucketName, objectName, time.Second*60, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
