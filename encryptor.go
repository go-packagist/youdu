package youdu

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
)

type Encryptor struct {
	appid  string
	aesKey string
	pkcs7  *Pkcs7
}

func NewEncryptor(appid, aesKey string) *Encryptor {
	return &Encryptor{
		appid:  appid,
		aesKey: aesKey,
		pkcs7:  NewPkcs7(),
	}
}

func (e *Encryptor) Encrypt(plaintext string) (string, error) {
	// decode key
	key, err := base64.StdEncoding.DecodeString(e.aesKey)
	if err != nil {
		return "", err
	}

	// plainText
	plainText := make([]byte, 0)

	randBs := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, randBs)
	if err != nil {
		return "", err
	}

	lenBs := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBs, uint32(len([]byte(plaintext))))

	plainText = append(plainText, randBs...)
	plainText = append(plainText, lenBs...)
	plainText = append(plainText, []byte(plaintext)...)
	plainText = append(plainText, []byte(e.appid)...)

	// encrypt
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	plainText = e.pkcs7.Padding(plainText)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (e *Encryptor) MustEncrypt(plaintext string) string {
	ciphertext, err := e.Encrypt(plaintext)

	if err != nil {
		panic(err)
	}

	return ciphertext
}

func (e *Encryptor) Decrypt(ciphertext string) (string, error) {
	// key
	key, err := base64.StdEncoding.DecodeString(e.aesKey)
	if err != nil {
		return "", err
	}

	// cipherText
	cipherText, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// len valid
	if len(cipherText)%len(key) != 0 {
		return "", errors.New("invalid ciphertext")
	}

	// aes decrypt
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = e.pkcs7.Unpadding(plainText)

	// rawMessage
	var length int32
	if err := binary.Read(bytes.NewBuffer(plainText[16:20]), binary.BigEndian, &length); err != nil {
		return "", err
	}
	if len(plainText) < int(20+length) {
		return "", errors.New("invalid ciphertext")
	}

	return string(plainText[20 : 20+length]), nil
}

func (e *Encryptor) MustDecrypt(ciphertext string) string {
	plaintext, err := e.Decrypt(ciphertext)

	if err != nil {
		panic(err)
	}

	return plaintext
}

type Pkcs7 struct {
	blockSize int
}

func NewPkcs7() *Pkcs7 {
	return &Pkcs7{
		blockSize: 32,
	}
}

func (p *Pkcs7) Padding(content []byte) []byte {
	padding := p.blockSize - (len(content) % p.blockSize)

	if padding == 0 {
		padding = p.blockSize
	}

	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(content, padtext...)
}

func (p *Pkcs7) Unpadding(content []byte) []byte {
	if len(content) == 0 {
		return nil
	}

	padding := content[len(content)-1]
	if int(padding) > len(content) || int(padding) > p.blockSize {
		return nil
	} else if padding == 0 {
		return nil
	}

	for i := len(content) - 1; i > len(content)-int(padding)-1; i-- {
		if content[i] != padding {
			return nil
		}
	}

	return content[:len(content)-int(padding)]
}
