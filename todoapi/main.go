package main

import (

	"github.com/gin-gonic/gin"
	"todoapi/database"
	"todoapi/routers"
)

func main() {

	database.Connect()
	r := gin.Default()
	routers.Setup(r)
	r.Run() 
}
