package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Jason-cqtan/simple-blog/config"
	"github.com/Jason-cqtan/simple-blog/database"
	"github.com/Jason-cqtan/simple-blog/routes"
	"github.com/Jason-cqtan/simple-blog/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	seed := flag.Bool("seed", false, "initialize the database and load seed data, then exit")
	flag.Parse()

	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if *seed {
		if err := database.Seed(db); err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}
		return
	}

	router := gin.Default()

	tmpl, err := utils.LoadTemplates("./views")
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}
	router.SetHTMLTemplate(tmpl)
	router.Static("/static", "./static")

	routes.SetupRoutes(router, db, cfg)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
