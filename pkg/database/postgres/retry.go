package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// FunctionWithError is a type alias for a function that returns an error.
type FnDBConnection func() (*gorm.DB, error)

// Retry function takes retryCount and a function with an error return type.
func Retry(retryCount int, fn FnDBConnection) (*gorm.DB, error) {
	var err error

	for i := 1; i <= retryCount; i++ {

		db, err := fn()
		if err == nil {
			return db, nil
		}
		log.Printf("Attempt %d: Failed with error: %v", i, err)
		time.Sleep(time.Duration(i) * time.Second)
	}

	return nil, fmt.Errorf("all %d attempts failed, last error: %w", retryCount, err)
}
