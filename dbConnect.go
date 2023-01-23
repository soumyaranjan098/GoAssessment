package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

var Uri = "root:!@!@soumya098@tcp(localhost:3306)/mydb"

func DataMigrate() {
	db, err = gorm.Open(mysql.Open(Uri))
	if err != nil {
		fmt.Println(err.Error())
		panic("connection failed")
	}

	db.AutoMigrate(&employee{})
}
