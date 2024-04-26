package main

import (
	"bwastartup/app/database"
	"bwastartup/app/entity"
	"bwastartup/app/model"
	"bwastartup/dto"
	"context"
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func main() {
	database.Connect()

	//standart gorm
	var users []dto.Users
	database.DB.Find(&users)
	for _, list := range users {
		fmt.Println(list.Name.String)

	}

	//init database using in gorm generate
	u := model.Use(database.DB)
	//using DAO User
	listUser, err := u.User.WithContext(context.Background()).Find()
	//fmt.Println(listUser.Name)
	for _, list := range listUser {
		fmt.Println(list.Name)
	}

	user, err := u.User.Find()
	for _, list := range user {
		fmt.Println(list.Name)
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
	fmt.Println("hasil result ", *result)

	//database.DB.Model(&users).Clauses(clause.Returning{}).Where("role = ?", "admin").Update("salary", gorm.Expr("salary * ?", 2))

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

	result_delete, err := u.User.Where(u.User.ID.Eq(46)).Delete()
	//disini walau id nya sebenarnya sudah dihapus tidak mengeluarkan error berbeda dengan fungsi update yang akan mengeluarkan error jika id tidak ketemu
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result_delete.Error)
	//end dao user

}
