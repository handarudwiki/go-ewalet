package component

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/ewalet?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return db
}
