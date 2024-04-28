package main

import (
	"bwastartup/domain/helper/mysql"
	"bwastartup/domain/helper/redis"
	v1 "bwastartup/routes/v1"
	v2 "bwastartup/routes/v2"
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/didip/tollbooth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	mysql.ConnectDatabase()
	mysql.AutoMigrateDB() //make sure success

	redis.SetupRedis()
	redis.SetupCacheChannel()

	router := gin.Default()
	router.Use(helmet.Default())
	router.Use(cors.Default())
	limiter := tollbooth.NewLimiter(100, nil) //global limitter

	v1.Setup(router, limiter)
	v2.Setup(router)

	var port = ":" + os.Getenv("PORT")
	fmt.Println("ini port golang ", port)
	err := router.Run(port)
	if err != nil {
		fmt.Println("Error running API in port", err)
	}
}
