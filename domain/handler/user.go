package handler

import (
	"bwastartup/domain/auth"
	"bwastartup/domain/helper"
	"bwastartup/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		messageError := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", messageError)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	registerUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	generateToken, err := h.authService.GenerateToken(registerUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(registerUser, generateToken)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		messageError := gin.H{"errors": errors}
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	loginUser, err := h.userService.Login(input)
	if err != nil {
		messageError := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	generateToken, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		messageError := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loginUser, generateToken)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
