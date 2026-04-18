package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	DatabaseURL string
	Port string
	StaleBundleAge time.Duration
	AllowedOrigins []string
}

func getEnvOrDefault(key , fallback string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return fallback
}

func Load() ( *Config , error ) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil , fmt.Errorf("DATABASE_URL is required")
	}

	port := getEnvOrDefault("PORT" , "3000")

	staleBundleAgeStr := getEnvOrDefault("STALE_BUNDLE_AGE", "2160h")

	staleBundleAge , err := time.ParseDuration(staleBundleAgeStr)
	if err != nil {
		return nil , fmt.Errorf("invalid STALE_BUNDLE_AGE %q %w", staleBundleAgeStr , err)
	}

	originsStr := getEnvOrDefault( "ALLOWED_ORIGINS" , "http://localhost:5173")

	origins := strings.Split(originsStr , ",")
	if len(origins) == 1 && origins[0] == "" {
		return nil , fmt.Errorf("ALLOWED_ORIGINS cannot be empty")
	}

	return &Config{
		DatabaseURL: dbURL,
		Port: port,
		StaleBundleAge: staleBundleAge,
		AllowedOrigins: origins,
	} , nil
}