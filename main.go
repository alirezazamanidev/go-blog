package main

import (
	"fmt"
	"log"

	"github.com/alirezazamanidev/go-blog/app/common/middlewares"
	"github.com/alirezazamanidev/go-blog/app/configs"
	"github.com/alirezazamanidev/go-blog/app/routes"
	"github.com/alirezazamanidev/go-blog/db"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middlewares.ErrorHandler())
	config := configs.Load()
	database := db.SetUp(config)
	rdb := db.SetUpRedis(config)
	routes.SetUp(router, database, rdb)

	log.Printf("🚀 Server is running on http://localhost:%s", config.AppPort)
	if err := router.Run(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}

}
