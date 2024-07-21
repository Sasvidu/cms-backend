package controllers

import (
	"context"
	"net/http"

	"firebase.google.com/go/v4/auth"
	initializers "github.com/Sasvidu/cms-backend/Initializers"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type DeleteUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client, err := initializers.FirebaseApp.Auth(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting Auth client"})
		return
	}

	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password)

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":   u.UID,
		"email": u.Email,
	})
}

func DeleteUserByEmail(c *gin.Context) {
	var req DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client, err := initializers.FirebaseApp.Auth(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting Auth client"})
		return
	}

	userRecord, err := client.GetUserByEmail(ctx, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = client.DeleteUser(ctx, userRecord.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted user"})
}