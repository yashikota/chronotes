package minio

import (
	"bytes"

	"github.com/minio/minio-go/v7"
)

func SaveObject(file *bytes.Buffer, filename string, fileSize int64, mimeType string) (minio.UploadInfo, error) {
	uploadInfo, err := MinioClient.PutObject(MinioCtx, bucketName, filename, file, fileSize, minio.PutObjectOptions{ContentType: mimeType})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return uploadInfo, nil
}
