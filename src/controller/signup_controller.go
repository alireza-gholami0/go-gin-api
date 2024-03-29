package controller

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignupController struct {
	UserService services.UserService
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request models.SignupRequest
	var err error

	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, _ := sc.UserService.CreateUser(c, request)
	if response != nil {
		c.JSON(http.StatusOK, response)
		return
	}
}
