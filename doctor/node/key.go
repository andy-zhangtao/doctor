package node

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	zr "github.com/andy-zhangtao/gogather/random"
	"golang.org/x/crypto/ssh"
)

// generateSSHKey 生成SSH 私钥文件和公钥文件
func generateSSHKey() (filename string, key []byte, err error) {
	filename = zr.GetString(12)

	savePrivateFileTo := fmt.Sprintf("/key/%s.pem", filename)
	savePublicFileTo := fmt.Sprintf("/key/%s.pub", filename)
	bitSize := 4096

	privateKey, err := generatePrivateKey(bitSize)
	if err != nil {
		return filename, key, fmt.Errorf("GeneratePrivateKey Error: [%s]", err.Error())
	}

	publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		return filename, key, fmt.Errorf("GeneratePublicKey Error: [%s]", err.Error())
	}

	privateKeyBytes := encodePrivateKeyToPEM(privateKey)

	err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		return filename, key, fmt.Errorf("Write Private Key Error: [%s]", err.Error())
	}

	err = writeKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		return filename, key, fmt.Errorf("Write Public Key Error: [%s]", err.Error())
	}

	return
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)
	return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	return nil
}
