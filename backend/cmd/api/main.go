package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lbringer-code/oneLink/backend/internal/config"
	"github.com/Lbringer-code/oneLink/backend/internal/db"
	"github.com/Lbringer-code/oneLink/backend/internal/handler"
	"github.com/Lbringer-code/oneLink/backend/internal/repository"
	"github.com/Lbringer-code/oneLink/backend/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout , nil))
	slog.SetDefault(logger)

	cfg , err := config.Load()
	if err != nil {
		logger.Error("config failed to load" , "error" , err)
		os.Exit(1)
	}

	database , err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Error("database failed to connect" , "error" , err)
		os.Exit(1)
	}
	defer database.Close()
	logger.Info("database connected")

	err = db.RunMigrations(database , "sql/migrations" , logger)
	if err != nil {
		logger.Error("migrations failed to run" , "error" , err)
		os.Exit(1)
	}
	logger.Info("migrations applied")

	repo := repository.New(database)
	svc := service.New(repo , logger)
	h := handler.New(svc , logger , cfg.AllowedOrigins)

	srv := &http.Server {
		Addr: ":" + cfg.Port,
		Handler: h.Router(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	serverErr := make(chan error , 1)
	go func(){
		logger.Info("server starting" , "port" , cfg.Port)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err , http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	quit := make(chan os.Signal , 1)
	signal.Notify(quit , syscall.SIGINT , syscall.SIGTERM)
	
	select {
	case err := <- serverErr:
		logger.Error("server error" , "error" , err)
		os.Exit(1)
	case sig := <- quit:
		logger.Info("shutdown signal received" , "signal" , sig.String())
	}

	shutdownCtx , cancel := context.WithTimeout(context.Background() , 15 * time.Second)
	defer cancel()

	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error("graceful shutdown failed" , "error" , err)
	}

	logger.Info("server shutdown cleanly")
}