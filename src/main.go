package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startup() {
	// ValidateLocalEnv()
	DataBase()
	SeedData()
}

func app_itself() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/videos/list", GetAllVideos)
	router.GET("/api/videos/:id", GetVideoByID)
	router.Run("0.0.0.0:3005")
}

func main() {
	startup()
	app_itself()
}
