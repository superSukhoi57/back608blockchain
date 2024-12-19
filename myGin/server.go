package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobackend/myGin/controller"
	"gobackend/myGin/gorm/DBLink"
	_ "gobackend/myGin/gorm/DBLink"
	"os"
)

func main() {

	// 获取当前工作路径
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 打印当前工作路径
	fmt.Println("当前工作路径:", wd)
	//TODO：（和java一样）当前工作路径就是go.mod所在的路径！！

	//尝试获取数据库连接
	db := DBLink.GetDB()
	if db != nil {
		fmt.Println("成功连接数据库！")
	}

	fmt.Println("开始启动服务器！😘")
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"congratulations": "恭喜😘",
			"msg":             "你已经成功启动gin",
		})
	})

	controller.MyTestRoute(server)

	server.Run(":9090")

}
