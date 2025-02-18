package utils

import (
	"fmt"
	"github.com/sony/sonyflake"
	"log"
)

func GetID() string {
	// Create a new Sonyflake
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if flake == nil {
		log.Println("Sonyflake not created")
		return ""
	}

	id, err := flake.NextID()
	if err != nil {
		log.Println("Failed to generate ID")
		return ""
	}

	log.Printf("Generated ID By 雪花算法: %d\n", id)

	return fmt.Sprintf("%d", id)
}
