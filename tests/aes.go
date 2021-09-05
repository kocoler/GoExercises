package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	//"encoding/base64"
	//"encoding/hex"
	"fmt"
)

//
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//fmt.Println("true", blockSize, len(ciphertext), padding)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//fmt.Println("padding text true", padtext)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AesEncrypt 加密函数
func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	//fmt.Println("iv true", iv)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	//fmt.Println("after padding true", plaintext)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// encrypt str with AES256-CBC, padding with PKCS7
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	ivT := make([]byte, aes.BlockSize+len(plaintext))
	// initialization vector
	iv := ivT[:aes.BlockSize]
	//fmt.Println("ivv", iv)
	// block
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	blockSize := block.BlockSize()

	// PKCS7 padding
	padding := blockSize - len(plaintext)%blockSize
	//fmt.Println("false", blockSize, len(plaintext), padding)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//fmt.Println("padding text", padtext)
	plaintext = append(plaintext, padtext...)
	//fmt.Println("after padding", plaintext)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)

	return (crypted), nil
}


func DecodeAES(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	iv := make([]byte, blockSize)
	blockMode := cipher.NewCBCDecrypter(block, iv)

	origData := make([]byte, len(data))
	blockMode.CryptBlocks(origData, data)
	length := len(origData)
	unpadding := int(origData[length-1])

	return origData[:(length - unpadding)], nil
}


func main() {
	key := []byte("wwwwwwwwwwwwwwww")
	//plaintext := []byte("hello ming")
	plaintext := []byte{123,34,109,101,115,115,97,103,101,95,116,121,112,101,34,58,49,44,34,118,101,114,115,105,111,110,34,58,34,53,49,34,44,34,115,116,97,116,117,115,34,58,50,44,34,99,97,117,115,101,34,58,34,100,101,112,108,111,121,32,115,117,99,99,101,115,115,34,125}

	c := make([]byte, aes.BlockSize+len(plaintext))
	iv := c[:aes.BlockSize]

	//加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	//打印加密base64后密码
	fmt.Println(ciphertext)

	//ciphertext := []byte{237,117,59,235,152,188,180,203,199,0,159,109,190,189,141,92,220,99,36,8,161,138,17,231,146,126,107,190,5,117,126,75,156,169,136,234,220,148,92,191,179,100,222,252,141,194,43,117,143,210,163,230,213,43,70,200,46,164,220,255,1,58,194,156,97,52,7,112,43,32,149,209,105,127,38,215,195,46,253,104,230,228,168,105,196,105,153,130,134,230,183,166,133,31,237,10}
	ciphertext, err = encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}


	fmt.Println(ciphertext)
	//fmt.Println(string(plaintext))
}
