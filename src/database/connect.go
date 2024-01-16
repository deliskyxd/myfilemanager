package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() {
    DNS := "root:@tcp(localhost:3307)/gofilemanager?charset=utf8&parseTime=True&loc=Local"

    database, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    DB = database
}
