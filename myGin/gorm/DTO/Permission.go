package DTO

import "time"

type Permission struct {
	Id          int       `json:"id"`
	Create_time time.Time `json:"create___time"`
	File_hash   string    `json:"file_Hash"`
}
