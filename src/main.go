package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	DataBase()
	router := gin.Default()
	router.Use(cors.Default())
	router.Run("0.0.0.0:3005")
}
