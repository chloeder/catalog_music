package main

import (
	"catalog-music/internal/configs"
	"catalog-music/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a new Gin router with default middleware
	r := gin.Default()

	// Declare config variable to store application configuration
	var (
		cfg *configs.Config
	)

	// Initialize configuration with specified options:
	// - Set config folder path
	// - Set config file name
	// - Set config file type
	if err := configs.Init(
		configs.WithConfigFolders([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	); err != nil {
		// If configuration initialization fails, log the error and exit
		log.Fatalf("failed to initialize configs: %v", err)
	}

	// Get the configuration instance
	cfg = configs.GetConfig()

	// Initialize the database connection
	_, err := internalsql.Connect(cfg.Database.DataSourcesName)
	if err != nil {
		// If database connection fails, log the error and exit
		log.Fatalf("failed to connect to the database: %v", err)
	}

	r.Run(cfg.Service.Port)
}
