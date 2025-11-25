package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database       DatabaseConfig
	Server         ServerConfig
	JWT            JWTConfig
	Storage        StorageConfig
	Redis          RedisConfig
	App            AppConfig
	MigrationsPath string
	AutoMigrate    bool
}

type DatabaseConfig struct {
	URL            string
	Host           string
	Port           string
	Name           string
	User           string
	Password       string
	SSLMode        string
	MaxConnections int
	MaxIdleConns   int
	MaxLifetime    time.Duration
}

type ServerConfig struct {
	Port           string
	Host           string
	Environment    string
	CORSOrigins    []string
	TrustedProxies []string
	LogLevel       string
	LogFormat      string
}

type JWTConfig struct {
	Secret                 string
	ExpiresIn              time.Duration
	RefreshTokenExpiresIn  time.Duration
	RefreshTokenCookieName string
}

type StorageConfig struct {
	Provider     string
	LocalPath    string
	AWSRegion    string
	AWSBucket    string
	AWSAccessKey string
	AWSSecretKey string
	AWSEndpoint  string
}

type RedisConfig struct {
	URL      string
	Password string
	CacheTTL time.Duration
}

type AppConfig struct {
	URL         string
	FrontendURL string
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load(".env.development")

	return &Config{
		Database: DatabaseConfig{
			URL:            getEnv("DATABASE_URL", "postgres://postgres:password@db:5432/rental_mvp?sslmode=disable"),
			Host:           getEnv("DATABASE_HOST", "db"),
			Port:           getEnv("DATABASE_PORT", "5432"),
			Name:           getEnv("DATABASE_NAME", "rental_mvp"),
			User:           getEnv("DATABASE_USER", "postgres"),
			Password:       getEnv("DATABASE_PASSWORD", "password"),
			SSLMode:        getEnv("DATABASE_SSL_MODE", "disable"),
			MaxConnections: getEnvAsInt("DATABASE_MAX_CONNECTIONS", 25),
			MaxIdleConns:   getEnvAsInt("DATABASE_MAX_IDLE_CONNECTIONS", 10),
			MaxLifetime:    time.Duration(getEnvAsInt("DATABASE_MAX_LIFETIME_MINUTES", 5)) * time.Minute,
		},
		Server: ServerConfig{
			Port:           getEnv("PORT", "8080"),
			Host:           getEnv("HOST", "0.0.0.0"),
			Environment:    getEnv("ENVIRONMENT", "development"),
			CORSOrigins:    strings.Split(getEnv("CORS_ORIGINS", "http://localhost:3000"), ","),
			TrustedProxies: strings.Split(getEnv("TRUSTED_PROXIES", "127.0.0.1"), ","),
			LogLevel:       getEnv("LOG_LEVEL", "info"),
			LogFormat:      getEnv("LOG_FORMAT", "json"),
		},
		JWT: JWTConfig{
			Secret:                 getEnv("JWT_SECRET", "your-secret-key"),
			ExpiresIn:              parseDuration(getEnv("JWT_EXPIRES_IN", "15m")),
			RefreshTokenExpiresIn:  parseDuration(getEnv("REFRESH_TOKEN_EXPIRES_IN", "168h")),
			RefreshTokenCookieName: getEnv("REFRESH_TOKEN_COOKIE_NAME", "refresh_token"),
		},
		Storage: StorageConfig{
			Provider:     getEnv("STORAGE_PROVIDER", "local"),
			LocalPath:    getEnv("LOCAL_STORAGE_PATH", "./uploads"),
			AWSRegion:    getEnv("AWS_REGION", "us-east-1"),
			AWSBucket:    getEnv("AWS_BUCKET", ""),
			AWSAccessKey: getEnv("AWS_ACCESS_KEY_ID", ""),
			AWSSecretKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
			AWSEndpoint:  getEnv("AWS_ENDPOINT", ""),
		},
		Redis: RedisConfig{
			URL:      getEnv("REDIS_URL", "redis://localhost:6379/0"),
			Password: getEnv("REDIS_PASSWORD", ""),
			CacheTTL: time.Duration(getEnvAsInt("CACHE_TTL_MINUTES", 60)) * time.Minute,
		},
		App: AppConfig{
			URL:         getEnv("APP_URL", "http://localhost:3000"),
			FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
		},
		MigrationsPath: getEnv("MIGRATIONS_PATH", "./migrations"),
		AutoMigrate:    getEnvAsBool("AUTO_MIGRATE", true),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func parseDuration(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return 15 * time.Minute // default fallback
	}
	return duration
}
