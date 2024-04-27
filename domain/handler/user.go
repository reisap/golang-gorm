package handler

import (
	"bwastartup/domain/helper"
	"bwastartup/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		messageError := helper.FormatValidationError(err)
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	registerUser, err := h.userService.RegisterUser(input)
	if err != nil {
		messageError := helper.FormatValidationError(err)
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(registerUser, "tokentoken")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
