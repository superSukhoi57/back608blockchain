package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobackend/myGin/gorm/DBLink"
	"gobackend/myGin/gorm/DTO"
	"gobackend/myGin/utils"
	"net/http"
	"time"
)

func MyTestRoute(r *gin.Engine) {
	mytest := r.Group("/test")
	{
		mytest.GET("/test", test)
		mytest.POST("/encryption", encryptionTxt)
		mytest.POST("/upload", uploadMydata)
	}

}

// 测试服务器连通的
func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "test",
	})
}

// 可以返回请求中文件的哈希值
func encryptionTxt(c *gin.Context) {
	file, err := c.FormFile("files")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有文件"})
		return
	}

	//生成文件的哈希散列值
	hashString, ok := utils.FileSHA256(file)
	if !ok {
		c.JSON(http.StatusBadRequest, "文件加密失败")
	}

	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"content":  hashString,
	})

}

// 上传文件操作
func uploadMydata(c *gin.Context) {
	//取出请求里面的文件
	file, err := c.FormFile("files")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有文件"})
		return
	}

	//生成文件的哈希散列值
	hashString, ok := utils.FileSHA256(file)
	if !ok {
		c.JSON(http.StatusBadRequest, "文件加密失败")
	}
	fmt.Println("生成的散列值：" + hashString)

	//入库
	db := DBLink.GetDB()

	// 执行查询以获取所有表名
	var tables []string
	db.Raw("SHOW TABLES").Scan(&tables)
	// 打印所有表名
	for _, table := range tables {
		fmt.Println(table)
	}

	//先验证数据库文件是否唯一
	var count int64
	db.Table("files").Model(&DTO.File{}).Where("file_hash = ?", hashString).Count(&count)
	if count > 0 {
		fmt.Println("记录已存在")
		c.JSON(500, "记录已经存在")
		return
	}
	myfile := DTO.File{
		File_hash:   hashString,
		File_shares: 2,
		//Data这个值根据文件需要确定，是序列化的值还是文件描述（文件存在特定内存，然后是这个文件的描述）
		Data:       "example data",
		Race:       1,
		Age:        30,
		Blood_type: 1,
		Gender:     true,
		Height:     180.5,
		Weight:     75.0,
		Smk_stat:   0,
		Alc_stat:   0,
		//Create_time: time.Date(2024, 5, 6, 0, 0, 0, 0, time.UTC),
		Create_time: time.Now(),
		Update_time: time.Now(),
	}
	//result := db.Table("files").Save(&myfile)//这tm是update的！
	result := db.Create(&myfile).Table("files")
	if result.Error != nil {
		fmt.Printf("数据保存失败: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据保存失败"})
		return
	}
	c.JSON(http.StatusOK, gin.Context{Accepted: []string{"数据保存成功！", hashString}})
}

// NumberAndRevenue 选择要铸造的份数以及确定将收取的创作者收益百分比接口
func NumberAndRevenue(c *gin.Context) {
	//路径参数是/numrev/:num/:rev
	number := c.Param("num")
	revenue := c.Param("rev")

	//其他逻辑……crud

	//调用sodility……
	c.JSON(200, gin.H{
		"number":  number,
		"revenue": revenue,
	})
}
