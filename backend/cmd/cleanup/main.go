package cleanup

import (
	"log/slog"
	"os"
	"time"

	"github.com/Lbringer-code/oneLink/backend/internal/config"
	"github.com/Lbringer-code/oneLink/backend/internal/db"
	"github.com/Lbringer-code/oneLink/backend/internal/repository"
	"github.com/Lbringer-code/oneLink/backend/internal/service"
	"github.com/joho/godotenv"
)


func main() {
	_ = godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout , nil))

	cfg , err := config.Load()
	if err != nil {
		logger.Error("config load failed" , "error" , err)
		os.Exit(1)
	}

	database , err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Error("db connect failed" , "error" , err)
		os.Exit(1)
	}

	repo := repository.New(database)
	svc := service.New(repo , logger)

	cutoff := time.Now().Add(-cfg.StaleBundleAge)
	logger.Info("running cleanup" , "cutoff" , cutoff , "max_age" , cfg.StaleBundleAge)

	count , err := svc.CleanupStaleBundles(cutoff)
	if err != nil {
		logger.Error("cleanup" , "error" , err)
		os.Exit(1)
	}

	logger.Info("cleaup completed" , "deleted_count" , count)
}