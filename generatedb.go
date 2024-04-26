package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./entity/repository",
		//WithUnitTest:   true,
		FieldNullable:  true,
		FieldCoverable: true,
		//Mode:           gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("connection database success !!")
	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		g.GenerateModel("users"),
		g.GenerateModel("transactions"),
		g.GenerateModel("campaigns"),
		g.GenerateModel("campaign_images"),
	)
	g.Execute()
}
