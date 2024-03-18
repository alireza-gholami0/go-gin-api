package controller

import (
	"github.com/alireza-gholami0/go-gin-api/src/models"
	"github.com/alireza-gholami0/go-gin-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	UserService services.UserService
}

func (lc *LoginController) Login(c *gin.Context) {
	var request models.LoginRequest
	var err error

	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	response, err := lc.UserService.Login(c, request)
	if response != nil {
		c.JSON(http.StatusOK, response)
		return
	}

}
