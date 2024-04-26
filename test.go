package main

import (
	"bwastartup/app/entity"
	"bwastartup/app/repository"
	"bwastartup/dto"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func main() {
	dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("connection database success !!")

	var users []dto.Users
	db.Find(&users)
	for _, list := range users {
		fmt.Println(list.Name.String)
		fmt.Println(list.Email.String)
	}

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

	where_user, err := u.User.Where(u.User.ID.Eq(1)).First()
	fmt.Println(where_user.Name)

	params_insert := entity.User{

		Name:           "ucupaditbakso",
		Occupation:     "it",
		Email:          "ucup@test.com",
		PasswordHash:   "test",
		AvatarFileName: "test.png",
		Role:           "admin",
		Token:          "test",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	err = u.User.Clauses(clause.Returning{}).Save(&params_insert)
	result, err := u.User.Last()
	fmt.Println(result.Name)

	//db.Model(&users).Clauses(clause.Returning{}).Where("role = ?", "admin").Update("salary", gorm.Expr("salary * ?", 2))

	params_update := entity.User{Name: "budi"}
	update_user, err := u.User.Where(u.User.ID.Eq(1)).Updates(&params_update)
	//update_user, err := u.User.Clauses(clause.Returning{}).Where(u.User.ID.Eq(1)).Update(u.User.Name, "joko")
	if err != nil {
		log.Fatal(err.Error())
	}
	result_update, err := u.User.Where(u.User.ID.Eq(1)).First()
	fmt.Println(update_user)
	fmt.Println(&params_update)
	fmt.Println(*result_update)

	//end dao user

}
