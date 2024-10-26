package minio

import (
	"log/slog"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func Connect() {
	endpoint := "minio:9000"
	accessKeyID := "root"
	secretAccessKey := os.Getenv("MINIO_PASSWORD")
	useSSL := false

	// Initialize minio client object.
	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Connected to Minio")
}
