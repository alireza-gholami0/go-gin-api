package routes

import (
	"github.com/alireza-gholami0/go-gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func Setup(time time.Duration, db gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewLoginRouter(time, db, publicRouter)
	NewSignupRouter(time, db, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middlewares.JwtAuthMiddleware("secret-key"))
	NewTaskRouter(time, db, protectedRouter)
}
