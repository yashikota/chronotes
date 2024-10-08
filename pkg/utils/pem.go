package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log/slog"
	"os"
)

var privateKey *ecdsa.PrivateKey

func LoadPrivateKeyFromEnv() {
	pemKey := os.Getenv("ECDSA_PRIVATE_KEY")
	if pemKey == "" {
		slog.Error("ECDSA_PRIVATE_KEY is not set")
		panic(errors.New("ECDSA_PRIVATE_KEY is not set"))
	}

	// Decode PEM
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil || block.Type != "EC PRIVATE KEY" {
		slog.Error("Failed to decode PEM block containing EC private key")
		panic(errors.New("failed to decode PEM block containing EC private key"))
	}

	// Parse EC private key
	var err error
	privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		slog.Error("Failed to parse EC private key")
		panic(err)
	}
}
