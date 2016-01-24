package crypt

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"hash"
	"log"
)

// EncryptFile encrypts file and returns encrypted file
func EncryptFile(file []byte, publicKey *rsa.PublicKey) ([]byte){
	var encrypted, label []byte

	encrypted = encryptOaep(publicKey, file, label)

	//fmt.Printf("OAEP Encrypted [%s] to \n[%x]\n", string(file), encrypted)

	return encrypted
}

// DecryptFile decrypts file and returns OG file
func DecryptFile(file []byte, privateKey *rsa.PrivateKey) ([]byte){
	var decrypted, label []byte

	decrypted = decryptOaep(privateKey, file, label)

	//fmt.Printf("OAEP Decrypted [%x] to \n[%s]\n", file, decrypted)

	return decrypted
}

// GenerateKeys creates a public and private key for the current transfer
// returns both
func GenerateKeys() (*rsa.PublicKey, *rsa.PrivateKey) {
	var publicKey *rsa.PublicKey
	var privateKey *rsa.PrivateKey
	var err error

	// Generate Private Key
	if privateKey, err = rsa.GenerateKey(rand.Reader, 1024); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Private Key: %v\n\n\n", privateKey)

	// Validate Private Key -- Sanity checks on the key
	if err = privateKey.Validate(); err != nil {
		log.Fatal(err)
	}

	// Precompute some calculations speeds up private key operations in the future
	privateKey.Precompute()

	//Public key address (of an RSA key)
	publicKey = &privateKey.PublicKey

	return publicKey, privateKey
}

//OAEP Encrypt
func encryptOaep(publicKey *rsa.PublicKey, plainText, label []byte) (encrypted []byte) {
	var err error
	var md5Hash hash.Hash

	md5Hash = md5.New()

	if encrypted, err = rsa.EncryptOAEP(md5Hash, rand.Reader, publicKey, plainText, label); err != nil {
		log.Fatal(err)
	}

	return
}

func decryptOaep(privateKey *rsa.PrivateKey, encrypted, label []byte) (decrypted[]byte) {
	var err error
	var md5Hash hash.Hash

	md5Hash = md5.New()
	if decrypted, err = rsa.DecryptOAEP(md5Hash, rand.Reader, privateKey, encrypted, label); err != nil {
		log.Fatal(err)
	}

	return
}


