package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
)

// FileSHA256 接收一个*multipart.FileHeader 类型的指针作为参数
/*
*multipart.FileHeader是用于处理HTTP请求中的multipart/form-data类型数据的一部分。
当一个HTTP请求使用multipart/form-data编码时，它通常用于上传文件。
*/
func FileSHA256(file *multipart.FileHeader) (string, bool) {
	// Open the uploaded files
	f, err := file.Open()
	if err != nil {
		fmt.Println("无法打开文件！")
		return "", false
	}
	//“Unhandled error”，这通常意味着程序在尝试关闭一个资源（如文件、网络连接等）时遇到了错误，但是没有对这个错误进行处理。
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Printf("无法关闭文件: %v", err)
		}
	}(f)

	// 创建一个新的 SHA-256 哈希器
	hash := sha256.New()

	// 将文件内容复制到哈希器中
	if _, err := io.Copy(hash, f); err != nil {
		fmt.Println("无法计算哈希值:", err)
		return "", false
	}

	// 获取哈希值的字节数组
	hashInBytes := hash.Sum(nil)

	// 将字节数组转换为十六进制字符串
	hashString := hex.EncodeToString(hashInBytes)

	// 输出哈希值
	fmt.Println("FileSHA256计算得到文件的 SHA-256 哈希值:", hashString)
	return hashString, true
}
