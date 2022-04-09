package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
)

func Encrypt(input string) string {

	text := []byte(input)

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	return base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, text, nil))
}

func Decrypt(input string) string {
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Fatal("error:", err)
	}
	cypher := []byte(data)

	c, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(cypher) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := cypher[:nonceSize], cypher[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return string(plaintext)
}
