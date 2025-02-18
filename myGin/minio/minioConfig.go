package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var minioClient *minio.Client

func init() {
	url := viper.GetString("minio.host")
	accessKey := viper.GetString("minio.access_key")
	secretKey := viper.GetString("minio.secret_key")
	//fmt.Println("main的viper配置在这个包前生效！url:", url)
	fmt.Print("main的viper配置在这个包前生效！url:", url)
	fmt.Printf("accessKey:%s,secretKey:%s", accessKey, secretKey)
	var err error
	//不用：=
	minioClient, err = minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // 如果MinIO服务器使用HTTPS，请设置为true
	})
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}
}

func myminio() {
	// 确保在使用minioClient之前已经初始化
	if minioClient == nil {
		log.Fatal("MinIO client is not initialized")
	}

	// 检查桶是否存在。
	bucketName := "test"
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalf("Failed to check if bucket exists: %v", err)
	}

	// 如果桶不存在，则创建它。
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
		fmt.Printf("Bucket '%s' created successfully.\n", bucketName)
	} else {
		fmt.Printf("Bucket '%s' already exists.\n", bucketName)
	}

	// 遍历./pictures目录并将所有文件上传到桶。
	err = filepath.Walk("./pictures", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只上传文件，跳过目录。
		if info.IsDir() {
			return nil
		}

		// 上传文件到桶。
		objectName := filepath.Base(path)
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = minioClient.PutObject(context.Background(), bucketName, objectName, file, info.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
		if err != nil {
			return err
		}

		fmt.Printf("File '%s' uploaded successfully.\n", objectName)
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to upload files: %v", err)
	}
}
