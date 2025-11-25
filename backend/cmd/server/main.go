package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	apimiddleware "github.com/nyunja/30budget/backend/internal/api/middleware"
	"github.com/nyunja/30budget/backend/internal/api/routes"
	"github.com/nyunja/30budget/backend/internal/config"
	"github.com/nyunja/30budget/backend/internal/db" // Reverted to internal/db
	"github.com/nyunja/30budget/backend/internal/utils"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger := utils.NewLogger()
	defer logger.Sync()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Initialize database
	dbPool, err := db.NewConnection(cfg.Database) // Reverted to db.NewConnection
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer dbPool.Close()

	// Run migrations if enabled
	if cfg.AutoMigrate {
		if err := db.RunMigrations(cfg.Database.URL, cfg.MigrationsPath); err != nil { // Reverted to db.RunMigrations
			logger.Fatal("Failed to run migrations", zap.Error(err))
		}
		logger.Info("Database migrations completed successfully")
	}

	// Initialize router
	r := chi.NewRouter()

	// Add middleware in order
	r.Use(middleware.RequestID)
	r.Use(apimiddleware.NewLogging(logger)) // Use our custom zap-based logging middleware
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Configure CORS
	logger.Info("CORS configuration",
		zap.Strings("allowed_origins", cfg.Server.CORSOrigins))
	r.Use(apimiddleware.NewCORS(cfg.Server.CORSOrigins))

	// Add security headers after CORS to ensure they're applied to all responses
	r.Use(apimiddleware.SecurityHeaders)

	// Add request size limits (10MB max for all requests)
	r.Use(apimiddleware.RequestSizeLimit(10 * 1024 * 1024))

	// Health check endpoint
	r.Get("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","timestamp":"` + time.Now().UTC().Format(time.RFC3339) + `"}`))
	})

	// Initialize API routes
	routes.SetupRoutes(r, dbPool, cfg, logger)

	// Create server
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Starting server",
			zap.String("address", server.Addr),
			zap.String("host", cfg.Server.Host),
			zap.String("port", cfg.Server.Port),
			zap.String("environment", cfg.Server.Environment))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server stopped gracefully")
}
