package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

type Config struct {
	Host              string
	User              string
	Pass              string
	Port              string
	DBName            string
	WaitForConnection time.Duration
	RetryCount        int
	Scheme            string
}

type Option func(cfg *Config)

// New creates a new database connection.
// options: a list of options to apply to the connection.
// returns: a database connection.
func New(options ...func(cfg *Config)) *gorm.DB {

	cfg := &Config{}
	cfg.WaitForConnection = time.Second * 1
	cfg.RetryCount = 5
	cfg.Scheme = "public"

	for _, opt := range options {
		opt(cfg)
	}

	time.Sleep(cfg.WaitForConnection)

	fmt.Println("establishing connection to pg...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=%s",
		cfg.Host,
		cfg.User,
		cfg.Pass,
		cfg.DBName,
		cfg.Port,
		"disable",
		cfg.Scheme,
	)

	db, err := Retry(cfg.RetryCount, func() (*gorm.DB, error) {
		return gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	})

	if err != nil {
		panic(err)
	}

	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	return db

}

func WithSchema(schema string) Option {
	return func(cfg *Config) {
		cfg.Scheme = schema
	}
}

// WithHost returns a function that sets the specified host in the given Config.
func WithHost(host string) Option {
	return func(cfg *Config) {
		cfg.Host = host
	}
}

// WithUser returns a function that sets the specified user in the given Config.
func WithUser(user string) Option {
	return func(cfg *Config) {
		cfg.User = user
	}
}

// WithPass returns a function that sets the specified password in the given Config.
func WithDbName(dbName string) Option {
	return func(cfg *Config) {
		cfg.DBName = dbName
	}
}

// WithPass returns a function that sets the specified password in the given Config.
func WithPass(pass string) Option {
	return func(cfg *Config) {
		cfg.Pass = pass
	}
}

// WithPort returns a function that sets the specified port in the given Config.
// port: a string representing the port number.
// returns: an Option function that sets the given port in the Config.
func WithPort(port string) Option {
	return func(cfg *Config) {
		cfg.Port = port
	}
}

// WithWaitForConnection returns a function that sets the specified waitForConnection in the given Config.
func WithWaitForConnection(delay time.Duration) Option {
	return func(cfg *Config) {
		cfg.WaitForConnection = delay
	}
}

func WithRetryCount(retyCount int) Option {
	return func(cfg *Config) {
		cfg.RetryCount = retyCount
	}
}

// WithScheme returns a function that sets the specified scheme in the given Config.
func WithScheme(scheme string) Option {
	return func(cfg *Config) {
		cfg.Scheme = scheme
	}
}
