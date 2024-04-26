package main

import (
	"bwastartup/dto"
	"bwastartup/model/repository"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
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

	//for _, list := range users {
	//	fmt.Println(list.Name.String)
	//	fmt.Println(list.Email.String)
	//}

	//init database using in gorm generate
	u := repository.Use(db)
	//using DAO User
	listUser, err := u.User.WithContext(context.Background()).Find()
	//fmt.Println(listUser.Name)
	for _, list := range listUser {
		fmt.Println(list.Name)
		fmt.Println(list.Email)
	}

	user, err := u.User.Find()
	for _, list := range user {
		fmt.Println(list.Name)
		fmt.Println(list.Email)
	}

	oneuser, err := u.User.First()
	fmt.Println(oneuser.Name)
	//end dao user

}
