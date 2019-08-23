package utils

import (
	"crypto/rand"
	"crypto/rsa"
)

func RsaDecrypt(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaEncrypt(origData []byte, pub *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}
