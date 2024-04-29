package transactions

import (
	"bwastartup/src/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type transactionController struct {
	transactionService Service
}

func NewTransactionController(transactionService Service) *transactionController {
	return &transactionController{transactionService: transactionService}
}

func (h *transactionController) ListPaging(c *gin.Context) {
	limitQuery := c.Query("limit")
	pageQuery := c.Query("page")

	if limitQuery == "" || pageQuery == "" {
		user, err := h.transactionService.FindAllTbl()
		if err != nil {
			messageError := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Got Error when list user", http.StatusUnprocessableEntity, "error", messageError)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		formatter := gin.H{
			"data": user,
		}
		response := helper.APIResponse("List find all", http.StatusOK, "success", formatter)
		c.JSON(http.StatusOK, response)
	} else {
		limit, _ := strconv.Atoi(limitQuery)
		page, _ := strconv.Atoi(pageQuery)

		pagingUser, err := h.transactionService.FindPagingTbl(limit, page)
		if err != nil {
			messageError := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Got Error when list user", http.StatusUnprocessableEntity, "error", messageError)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		response := helper.APIResponse("List find with paging", http.StatusOK, "success", pagingUser)
		c.JSON(http.StatusOK, response)
	}

}
