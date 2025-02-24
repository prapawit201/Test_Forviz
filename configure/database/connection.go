package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("root:rootpassword@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=UTC")

	// เชื่อมต่อกับ MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to database error: %v", err)
		return nil, err
	}

	return db, nil
}
