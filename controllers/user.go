package controllers

import (
	"asignment/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUsers(ctx *gin.Context) {
	var input UserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	u := models.User{}
	u.FirstName = input.FirstName
	u.LastName = input.LastName
	u.Mobile = input.Mobile
	u.Email = input.Email

	_, err := u.SaveUser()
	if err != nil {
		fmt.Println("Not able to create User")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User Created successfully",
	})
}
