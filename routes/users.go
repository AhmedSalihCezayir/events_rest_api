package routes

import (
	"net/http"
	"strconv"

	"example.com/events-api/models"
	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil || user.Nickname == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User is created successfuly."})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID, user.IsAdmin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}

func getUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users. Try again later."})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func deleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id!"})
	}

	user, err := models.FindUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user. Try again later."})
	}

	err = user.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully."})
}
