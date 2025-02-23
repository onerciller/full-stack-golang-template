package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Lifecycle struct{}

func (c *Lifecycle) Provide(lc fx.Lifecycle, httpServer *httpserver.HttpServer, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting HTTP server")
			go func() {
				if err := httpServer.Start(); err != nil {
					log.Printf("HTTP server stopped with error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server")
			httpServer.Shutdown()
			db, err := db.DB()
			if err != nil {
				fmt.Println("Error getting database connection")
				return err
			}
			fmt.Println("Closing database connection")
			return db.Close()
		},
	})
}
