package main

import (
	"github.com/alireza-gholami0/go-gin-api/src/config"
	"github.com/alireza-gholami0/go-gin-api/src/routes"
	"github.com/alireza-gholami0/go-gin-api/src/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	utils.Loadenv()
	config.OpenDatabaseConnection()
	config.RabbitConnection()

	timeout := time.Duration(10) * time.Second

	engine := gin.Default()
	routes.Setup(timeout, *config.Database, engine)

	engine.Run(":8080")
	//if err != nil {
	//	panic(err)
	//}
}
