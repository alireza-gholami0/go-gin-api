package controller

import (
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	UserService services.UserService
}

func (sc *SignupController) Signup(c *gin.Context) {
	c.JSON(200, nil)
}
