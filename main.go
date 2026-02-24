package main

import (
	"fmt"
	"log"

	"github.com/Jason-cqtan/simple-blog/config"
	"github.com/Jason-cqtan/simple-blog/database"
	"github.com/Jason-cqtan/simple-blog/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")
	router.Static("/static", "./static")

	routes.SetupRoutes(router, db, cfg)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
