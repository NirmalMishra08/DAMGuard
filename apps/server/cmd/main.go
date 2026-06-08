package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/config"
	"main/internal/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("not able to load .env file")
		return
	}

	// routes logic

	router := gin.Default()
	routes.SetupRoutes(router, cfg)
	addr := fmt.Sprintf(":%d", cfg.PORT)

	defaultRouter(router, cfg)
	
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		fmt.Printf("Server running on port %s\n", cfg.PORT)
		if err := server.Start(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	// Accept SIGINT (Ctrl+C) and SIGTERM (docker stop)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("--- Shutting down server...----")

	// Graceful shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Close database connection
	sqlDB, err := db.DB()
	if err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println(" Database connection closed")
		}
	}

	log.Println(" Server exited gracefully")

}
