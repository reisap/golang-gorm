package main

import (
	"bwastartup/domain/helper/mysql"
	"bwastartup/domain/helper/redis"
	v1 "bwastartup/routes/v1"
	v2 "bwastartup/routes/v2"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	mysql.ConnectDatabase()
	mysql.AutoMigrateDB()
	redis.SetupRedis()
	redis.SetupCacheChannel()

	router := gin.Default()
	v1.Setup(router)
	v2.Setup(router)

	err := router.Run(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error running API in port", err)
	}
}
