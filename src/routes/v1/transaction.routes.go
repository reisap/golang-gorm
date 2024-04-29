package v1

import (
	"bwastartup/src/helper/mysql"
	"bwastartup/src/transactions"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(api *gin.RouterGroup) {
	repository := transactions.NewRepository(mysql.DB)
	service := transactions.NewService(repository)
	controller := transactions.NewTransactionController(service)

	api.GET("/transaction", controller.ListPaging)
}
