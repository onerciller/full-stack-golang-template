package module

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// AppLifecycleHooks contains application lifecycle hook functions
type AppLifecycleHooks struct {
	// Inject dependencies
	fx.In

	HttpServer *httpserver.HttpServer
	Database   *gorm.DB
}

// OnApplicationStart is called when the application has started
func (h *AppLifecycleHooks) OnApplicationStart(ctx context.Context) error {
	fmt.Println("Application has started successfully")

	// Log server details
	fmt.Println("Starting HTTP server...")

	// Start HTTP server in a goroutine
	go func() {
		if err := h.HttpServer.Start(); err != nil {
			log.Printf("HTTP server stopped with error: %v", err)
		}
	}()

	fmt.Println("HTTP server started successfully")
	return nil
}

// OnApplicationStop is called when the application is shutting down
func (h *AppLifecycleHooks) OnApplicationStop(ctx context.Context) error {
	fmt.Println("Application is shutting down gracefully")

	// Create a timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Step 1: Shutdown HTTP server with timeout
	fmt.Println("Stopping HTTP server...")
	if err := h.HttpServer.ShutdownWithTimeout(shutdownCtx); err != nil {
		fmt.Printf("Error shutting down HTTP server: %v\n", err)
		// Continue shutdown process despite error
	} else {
		fmt.Println("HTTP server stopped successfully")
	}

	// Step 2: Close database connection
	fmt.Println("Closing database connection...")
	db, err := h.Database.DB()
	if err != nil {
		fmt.Printf("Error getting database connection: %v\n", err)
		return err
	}

	// Close with timeout context
	closeErrChan := make(chan error, 1)
	go func() {
		closeErrChan <- db.Close()
	}()

	// Wait for database close to complete or timeout
	select {
	case err := <-closeErrChan:
		if err != nil {
			fmt.Printf("Error closing database connection: %v\n", err)
			return err
		}
		fmt.Println("Database connection closed successfully")
	case <-shutdownCtx.Done():
		if shutdownCtx.Err() == context.DeadlineExceeded {
			fmt.Println("WARNING: Database connection close timed out")
		}
		return shutdownCtx.Err()
	}

	fmt.Println("All resources have been cleaned up")
	return nil
}
