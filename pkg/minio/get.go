package minio

import (
	"time"
)

func GetObjectURL(objectName string) (string, error) {
	url, err := MinioClient.PresignedGetObject(MinioCtx, bucketName, objectName, time.Hour*24*7, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
