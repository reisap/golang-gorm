package main

import (
	"bwastartup/src/helper/mysql"
	"bwastartup/src/helper/redis"
	"bwastartup/src/routes/v1"
	"bwastartup/src/routes/v2"
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/didip/tollbooth"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
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

	//default framework using gin
	router := gin.Default()

	router.Use(helmet.Default())
	router.Use(cors.Default())
	router.Use(gzip.Gzip(gzip.BestSpeed))
	pprof.Register(router)

	limiter := tollbooth.NewLimiter(1000, nil) //global limitter

	v1.Setup(router, limiter)
	v2.Setup(router)

	var port = ":" + os.Getenv("PORT")
	fmt.Println("ini port golang ", port)
	err := router.Run(port)
	if err != nil {
		fmt.Println("Error running API in port", err)
	}
}
