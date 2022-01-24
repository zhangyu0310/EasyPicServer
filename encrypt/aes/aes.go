package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type Aes struct {
	// PrivateKey 加密秘钥, 长度分别是16,24,32位字符串,分别对应AES-128,AES-192,AES-256加密方式
	PrivateKey []byte
}

//pkcs4 填充模式
func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	//取余计算长度,判断加密的文本是不是blockSize的倍数,如果不是的话把多余的长度计算出来,用于补齐长度
	padding := blockSize - len(cipherText)%blockSize
	//补齐
	//Repeat: 把切片[]byte{byte(padding)}复制padding个然后合并成新的字节切片返回
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func (encryption *Aes) Encrypt(originData []byte) ([]byte, error) {
	//创建加密算法的实例
	block, err := aes.NewCipher(encryption.PrivateKey)
	if err != nil {
		return nil, err
	}

	//获取块的大小
	blockSize := block.BlockSize()

	//对数据进行填充,让数据的长度满足加密需求
	originData = pkcs7Padding(originData, blockSize)

	//采用aes加密方式中的CBC加密模式
	blockMode := cipher.NewCBCEncrypter(block, encryption.PrivateKey[:blockSize])
	encrypted := make([]byte, len(originData))

	//执行加密
	blockMode.CryptBlocks(encrypted, originData)

	return encrypted, nil
}

func (encryption *Aes) EncryptAndBase64(pwd []byte) (string, error) {
	//进行aes加密
	result, err := encryption.Encrypt(pwd)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(result), err
}

//填充的反向操作,删除填充的字符串
func pkcs7UnPadding(originData []byte) ([]byte, error) {
	//获取数据长度
	length := len(originData)
	if length <= 0 {
		return nil, errors.New("加密字符串长度不符合要求")
	}
	//获取填充字符串的长度
	unPadding := int(originData[length-1])
	//截取切片,删除填充的字节,并且返回明文
	return originData[:(length - unPadding)], nil
}

func (encryption *Aes) Decrypt(encrypted []byte) ([]byte, error) {
	//创建加密算法的实例
	block, err := aes.NewCipher(encryption.PrivateKey)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//创建加密实例
	blockMode := cipher.NewCBCDecrypter(block, encryption.PrivateKey[:blockSize])
	originData := make([]byte, len(encrypted))
	//该函数可也用来加密也可也用来解密
	blockMode.CryptBlocks(originData, encrypted)
	//取出填充的字符串
	originData, err = pkcs7UnPadding(originData)
	if err != nil {
		return nil, err
	}
	return originData, nil
}

func (encryption *Aes) DecryptFromBase64(pwd string) ([]byte, error) {
	pwdByte, err := base64.RawURLEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	return encryption.Decrypt(pwdByte)
}
