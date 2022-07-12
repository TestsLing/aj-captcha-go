package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

func AesEncrypt(data, key string) string {
	if key == "" {
		return data
	}
	return base64.StdEncoding.EncodeToString(AesEncryptToBytes([]byte(data), []byte(key)))
}

func AesEncryptToBytes(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS5Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func AesDecrypt(point string, key string) string {
	encryptBytes, _ := base64.StdEncoding.DecodeString(point)
	info := AESDecryptECB(encryptBytes, []byte(key))
	return string(info)
}

func AESDecryptECB(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS5UnPadding(decrypted)
}

// PKCS5UnPadding 删除pks5填充的尾部数据
func PKCS5UnPadding(origData []byte) []byte {
	// 1. 计算数据的总长度
	length := len(origData)
	// 2. 根据填充的字节值得到填充的次数
	number := int(origData[length-1])
	// 3. 将尾部填充的number个字节去掉
	return origData[:(length - number)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
