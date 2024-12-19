package DTO

import (
	"time"
)

// 对于File_name字段，可以使用gorm:"column:file_name"标签来指定数据库中的列名为file_name。
type File struct {
	File_name   string    `gorm:"column:file_name"`
	File_hash   string    `gorm:"column:file_hash"`
	File_shares int       `gorm:"column:File_shares"`
	Data        string    `gorm:"column:data"`
	Race        int       `gorm:"column:race"`
	Age         int       `gorm:"column:age"`
	Blood_type  int       `gorm:"column:blood_type"`
	Gender      bool      `gorm:"column:gender"`
	Height      float32   `gorm:"column:height"`
	Weight      float32   `gorm:"column:weight"`
	Smk_stat    int       `gorm:"column:smk_stat"`
	Alc_stat    int       `gorm:"column:alc_stat"`
	Other       string    `gorm:"column:other"`
	Description string    `gorm:"column:description"`
	Create_time time.Time `gorm:"column:create_time"`
	Update_time time.Time `gorm:"column:update_time"`
}
