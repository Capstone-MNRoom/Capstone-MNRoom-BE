package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"be9/mnroom/config"
	"be9/mnroom/factory"
	_middlewares "be9/mnroom/middlewares"
	"be9/mnroom/migration"
	"be9/mnroom/routes"
)

func main() {
	dbConn := config.InitDB()
	migration.InitMigrate(dbConn)
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	_middlewares.LogMiddlewares(e)

	// Start server in background goroutine
	go func() {
		if err := e.Start(":8000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("Error starting server: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	e.Logger.Info("Gracefully shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Server forced to shutdown: ", err)
	}

	// Close database connection
	if sqlDB, err := dbConn.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			e.Logger.Error("Error closing database connection: ", err)
		}
	}

	e.Logger.Info("Server exited cleanly")
}
