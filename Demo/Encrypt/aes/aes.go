/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2018/3/14 11:32
  */

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"encoding/base64"
	"bytes"
	"log"
)

/*
	D02goUwdclknLlesv5SErw==
	whatthefuck
*/

func testAesString() {
	ret, _ := AesEncrypt("whatthefuck")
	AesDecrypt(ret)
}

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("t66ycaoliushequ1")

	result, err := aesEncrypt([]byte("whatthefuck"), key)
	if err != nil {
		log.Println("aes Error:", err)
	}
	//fmt.Println("加密ret:", result)
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := aesDecrypt(result, key)
	log.Println("key1:", key)
	log.Println("decrypt1:", string(result))
	log.Println("decrypt1:", result)
	if err != nil {
		log.Println("aes Error:", err)
	}
	fmt.Println(string(origData))
}

func AesEncrypt(origData string) (string, error) {

	ori := []byte(origData)
	key := []byte("t66ycaoliushequ1")

	r, err := aesEncrypt(ori, key)
	if err != nil {
		log.Println("AES ENCRYPT ERROR:", err)
		return "", err
	}
	ret := base64.StdEncoding.EncodeToString(r)
	log.Println("加密后:", ret)
	return ret, nil
}

func AesDecrypt(crypted string) (string, error) {

	//cry := []byte(crypted)
	cry, _ := base64.StdEncoding.DecodeString(crypted)
	key := []byte("t66ycaoliushequ1")

	r, err := aesDecrypt(cry, key)
	if err != nil {
		log.Println("AES DECRYPT ERROR:", err)
		return "", err
	}
	ret := string(r)
	log.Println("解密后:", ret)
	return ret, nil
}


func aesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func aesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func main() {
	//testAes()
	testAesString()
}

