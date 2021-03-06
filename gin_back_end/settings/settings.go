package settings

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var STORAGE_DIR string

func init() {
	dsn := "root:woaini123@tcp(127.0.0.1:3306)/easyvideo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	STORAGE_DIR = "/Users/nocilantro/Downloads/storage"

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
