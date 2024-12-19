package DTO

import "time"

type User struct {
	Id          int       `json:"id"`
	Address     string    `json:"address"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Picture     string    `json:"picture"`
	Create_time time.Time `json:"create_Time"`
}
