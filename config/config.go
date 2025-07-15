package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
}

type AppConfig struct {
	Name        string
	Environment string
	Port        string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

type CORSConfig struct {
	AllowOrigins string
}

func loadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("warning: could not load env file: %s (%v)\n", envFile, err)
	} else {
		log.Printf("loaded environment: %s\n", envFile)
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}

func IsProduction() bool {
	return LoadConfig().App.Environment == "production"
}

func LoadConfig() *Config {
	if cfg != nil {
		return cfg
	}

	loadEnv()

	cfg = &Config{
		App: AppConfig{
			Name:        getEnv("APP_NAME", "light_pos"),
			Environment: getEnv("APP_ENV", "development"),
			Port:        getEnv("APP_PORT", "3000"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "light_pos"),
		},
		JWT: JWTConfig{
			AccessSecret:  getEnv("JWT_ACCESS_SECRET", "access-secret"),
			RefreshSecret: getEnv("JWT_REFRESH_SECRET", "refresh-secret"),
			AccessTTL:     time.Duration(getEnvInt("JWT_ACCESS_TTL", 15)) * time.Minute,
			RefreshTTL:    time.Duration(getEnvInt("JWT_REFRESH_TTL", 60)) * time.Minute,
		},
		CORS: CORSConfig{
			AllowOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000, http://localhost:5173"),
		},
	}

	return cfg
}
