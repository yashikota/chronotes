package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func generateECDSAKeys() (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECDSA private key: %v", err)
	}
	return privateKey, nil
}

func saveECDSAKeysToFile(privateKey *ecdsa.PrivateKey) error {
	dir := ".certs"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	privBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("unable to marshal ECDSA private key: %v", err)
	}

	privPem := pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes}
	keyOut, err := os.Create(".certs/ec_key.pem")
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer keyOut.Close()

	if err := pem.Encode(keyOut, &privPem); err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}

func main() {
	privateKey, err := generateECDSAKeys()
	if err != nil {
		log.Fatalf("Could not generate private key: %v", err)
	}

	if err := saveECDSAKeysToFile(privateKey); err != nil {
		log.Fatalf("Could not save private key: %v", err)
	}

	log.Println("ECDSA private key generated and saved successfully.")
}
