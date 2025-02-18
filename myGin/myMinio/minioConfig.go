package myMinio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
)

var minioClient *minio.Client

func init() {
	url := viper.GetString("minio.host") + ":" + viper.GetString("minio.port")
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

func GetMinioClient() *minio.Client {
	return minioClient
}
