package KitLib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

type Secret struct {
	_KEY string  "2iGVF52I" // 默认密钥，占8字节
}

// 默认密钥
//const _KEY string = "fY2o2igL" // 占8字节

// Encode 加密
func (kit *Secret) Encode(text string, key string) string {
	if len(key) != 8 {
		key = kit._KEY
	}
	defer func() { // 跳过致命错误使程序继续运行
		if e := recover(); e != nil { }
	}()
	cipherArr, err := SCEncryptString(text, key, "des")
	if err != nil {
		return ""
	}else {
		return cipherArr
	}
	// fmt.Printf("加密后字节数组：%v\n", cipherArr)
	// fmt.Printf("加密后16进制：%x\n", cipherArr)
}

// Decode 解密
func (kit *Secret) Decode(text string, key string) string {
	if len(key) != 8 {
		key = kit._KEY
	}
	defer func() { // 跳过致命错误使程序继续运行
		if e := recover(); e != nil { }
	}()
	originalBytes, err := SCDecryptString(text, key, "des")
	if err != nil {
		return ""
	}else{
		return originalBytes
	}
	// fmt.Println("解密后：", string(originalBytes))
}

// ======

// SCEncrypt DES加密
func SCEncrypt(originalBytes, key []byte, scType string) ([]byte, error) {
	// 1、实例化密码器block(参数为密钥)
	var err error
	var block cipher.Block
	switch scType {
	case "des":
		block, err = des.NewCipher(key)
	case "3des":
		block, err = des.NewTripleDESCipher(key)
	case "aes":
		block, err = aes.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	// fmt.Println("---blockSize---", blockSize)
	// 2、对明文填充字节(参数为原始字节切片和密码对象的区块个数)
	paddingBytes := PKCSSPadding(originalBytes, blockSize)
	// fmt.Println("填充后的字节切片：", paddingBytes)
	// 3、 实例化加密模式(参数为密码对象和密钥)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// fmt.Println("加密模式：", blockMode)
	// 4、对填充字节后的明文进行加密(参数为加密字节切片和填充字节切片)
	cipherBytes := make([]byte, len(paddingBytes))
	blockMode.CryptBlocks(cipherBytes, paddingBytes)
	return cipherBytes, nil
}

// SCDecrypt 解密字节切片，返回字节切片
func SCDecrypt(cipherBytes, key []byte, scType string) ([]byte, error) {
	// 1、实例化密码器block(参数为密钥)
	var err error
	var block cipher.Block
	switch scType {
	case "des":
		block, err = des.NewCipher(key)
	case "3des":
		block, err = des.NewTripleDESCipher(key)
	case "aes":
		block, err = aes.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	// 2、 实例化解密模式(参数为密码对象和密钥)
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	// fmt.Println("解密模式：", blockMode)
	// 3、对密文进行解密(参数为加密字节切片和填充字节切片)
	paddingBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(paddingBytes, cipherBytes)
	// 4、去除填充字节(参数为填充切片)
	originalBytes := PKCSSUnPadding(paddingBytes)
	return originalBytes, nil
}

// SCEncryptString SCEncryptString
func SCEncryptString(originalText, key, scType string) (string, error) {
	cipherArr, err := SCEncrypt([]byte(originalText), []byte(key), scType)
	if err != nil {
		panic(err)
	}
	base64str := base64.StdEncoding.EncodeToString(cipherArr)
	return base64str, nil
}

// SCDecryptString SCDecryptString
func SCDecryptString(cipherText, key, scType string) (string, error) {
	cipherArr, _ := base64.StdEncoding.DecodeString(cipherText)
	cipherArr, err := SCDecrypt(cipherArr, []byte(key), scType)
	if err != nil {
		panic(err)
	}
	return string(cipherArr), nil
}

// PKCSSPadding 填充字节的函数
func PKCSSPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	// fmt.Println("要填充的字节：", padding)
	// 初始化一个元素为padding的切片
	slice1 := []byte{byte(padding)}
	slice2 := bytes.Repeat(slice1, padding)
	return append(data, slice2...)
}

// ZeroPadding 填充字节的函数
func ZeroPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	// fmt.Println("要填充的字节：", padding)
	// 初始化一个元素为padding的切片
	slice1 := []byte{0}
	slice2 := bytes.Repeat(slice1, padding)
	return append(data, slice2...)
}

// PKCSSUnPadding 去除填充字节的函数
func PKCSSUnPadding(data []byte) []byte {
	unpadding := data[len(data)-1]
	result := data[:(len(data) - int(unpadding))]
	return result
}

// ZeroUnPadding 去除填充字节的函数
func ZeroUnPadding(data []byte) []byte {
	return bytes.TrimRightFunc(data, func(r rune) bool {
		return r == 0
	})
}

// ============================
