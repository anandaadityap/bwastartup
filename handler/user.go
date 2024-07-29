package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// Tangkap input dari user
	var input user.RegisterUserInput

	// map input dari user ke struct RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Register Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "")

	// struct diatas kita pasing sebagai parameter service
	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input email dan password
	// input ditangkap handler
	var input user.LoginInput
	// mapping input user ke input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// input struct di passing ke service
	// dicari dengan bantuan repository user dengan email x
	// mencocokan password
	loggedUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedUser, "")
	response := helper.ApiResponse("Login Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
