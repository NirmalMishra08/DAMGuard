package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/config"
	"main/internal/repository/clickhouse"
	"main/internal/repository/postgres/sqlc"

	"main/internal/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("not able to load .env file")
		return
	}

	pool, err := pgxpool.New(context.Background(), cfg.DB_URL)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("not able to connect postgres %v", err)
	}

	fmt.Print("conneted to postgres")

	defer pool.Close()

	store := sqlc.NewStore(pool)

	clickhouse_conn, err := clickhouse.New(cfg)
	if err != nil {
		 log.Fatal("not able to connect clickhouse: %v" , err)
	}


	// routes logic
	router := chi.NewRouter()
	routes.SetupRoutes(router, cfg, store, clickhouse_conn)
	addr := fmt.Sprintf(":%s", cfg.PORT)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		fmt.Printf("Server running on port %s\n", cfg.PORT)
		if err := srv.ListenAndServe(); err != nil {
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

	// Shutdown HTTP server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}

	// Close database connection
	pool.Close()
	log.Println("Database connection closed")
	log.Println("Server exited gracefully")

}
