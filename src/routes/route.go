package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func Setup(time time.Duration, db gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewSignupRouter(time, db, publicRouter)
}
