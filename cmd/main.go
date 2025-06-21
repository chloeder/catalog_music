package main

import (
	tracksHandler "catalog-music/external/handlers/tracks"
	tracksService "catalog-music/external/service/tracks"
	"catalog-music/external/spotify"
	"catalog-music/internal/configs"
	membershipsHandler "catalog-music/internal/handlers/memberships"
	"catalog-music/internal/models/memberships"
	membershipsRepo "catalog-music/internal/repositories/memberships"
	membershipsService "catalog-music/internal/services/memberships"
	"catalog-music/pkg/httpclient"
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
	db, err := internalsql.Connect(cfg.Database.DataSourcesName)
	if err != nil {
		// If database connection fails, log the error and exit
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Initialize the database tables
	db.AutoMigrate(&memberships.User{})

	// External services
	httpClient := httpclient.NewHTTPClient()
	spotifyOutbound := spotify.NewOutbound(cfg, httpClient)
	tracksService := tracksService.NewService(spotifyOutbound)
	tracksHandler := tracksHandler.NewHandler(r, tracksService)

	// Internal services
	membershipRepo := membershipsRepo.NewRepository(db)
	membershipService := membershipsService.NewService(cfg, membershipRepo)
	membershipHandler := membershipsHandler.NewHandler(r, membershipService)

	membershipHandler.AuthRoute()
	tracksHandler.TracksRoute()
	r.Run(cfg.Service.Port)
}
