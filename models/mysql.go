package models

import (
	"go-gorm/conf"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// 定义一个全局对象db
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(conf.Db["db1"].Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// db.SetMaxIdleConns(conf.Db["Db1"].MaxId)
	// db.SetMaxOpenConns(conf.Db["Db1"].MaxOpen)
	// db.SetConnMaxLifetime(2 * 3)
}

