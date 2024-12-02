package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToMysql() {
	dsn := "root:root@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to mysql: %v", err)
	}

	log.Println("Database initialized", DB)
}
