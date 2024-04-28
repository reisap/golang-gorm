package controller

import (
	"bwastartup/src/auth"
	"bwastartup/src/helper"
	"bwastartup/src/user"
	"fmt"
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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		messageError := gin.H{"errors": err.Error()}
		//response := helper.APIResponse("Check Email failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, messageError)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		messageError := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Check Email failed", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	var metaMessage string
	if isEmailAvailable == true {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been register"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	//tidak perlu bikin struct karena berupa form file
	file, err := c.FormFile("avatar")
	if err != nil {
		errors := helper.FormatValidationError(err)
		messageError := gin.H{"errors": errors}
		response := helper.APIResponse("Avatar filename not found in form", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//userId seharusnya dapat dari jwt
	userId := 1
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errors := helper.FormatValidationError(err)
		messageError := gin.H{"errors": errors}
		response := helper.APIResponse("Cannot upload avatar into directory", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.SaveAvatarUser(userId, path)
	if err != nil {
		errors := helper.FormatValidationError(err)
		messageError := gin.H{"errors": errors}
		response := helper.APIResponse("Cannot upload avatar", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar success uploaded", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
