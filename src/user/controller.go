package user

import (
	"bwastartup/src/auth"
	"bwastartup/src/helper"
	"bwastartup/src/user/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	userService Service
	authService auth.Service
}

func NewUserController(userService Service, authService auth.Service) *userController {
	return &userController{userService: userService, authService: authService}
}

func (h *userController) ListUserPaging(c *gin.Context) {
	user, _ := h.userService.UserPaging(1, 10)
	c.JSON(http.StatusOK, user)

}

func (h *userController) RegisterUser(c *gin.Context) {
	var input dto.RegisterUserInput
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
	formatter := dto.FormatUser(registerUser, generateToken)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var input dto.LoginUserInput
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
	formatter := dto.FormatUser(loginUser, generateToken)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userController) CheckEmailAvailability(c *gin.Context) {
	var input dto.CheckEmailInput
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
	if isEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been register"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userController) UploadAvatar(c *gin.Context) {
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
	userId := 2
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		messageError := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Cannot upload avatar into directory", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.SaveAvatarUser(userId, path)
	if err != nil {
		messageError := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Cannot upload avatar", http.StatusBadRequest, "error", messageError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar success uploaded", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
