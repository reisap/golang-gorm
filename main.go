package main

import (
	"bwastartup/dto"
	"bwastartup/query"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("connection database success !!")

	var users []dto.Users
	length := len(users)
	fmt.Println(length)
	db.Find(&users)
	length = len(users)
	fmt.Println(length)

	var listUsers []query.IUserDo
	db.Find(&listUsers)

}
