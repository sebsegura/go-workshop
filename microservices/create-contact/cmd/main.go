package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"seb7887/create-contact/internal/config"
	"seb7887/create-contact/internal/logger"
	"seb7887/create-contact/pkg/server"
	"syscall"
	"time"
)

func main() {
	var (
		serverPort = config.GetConfig().Port
		httpAddr = fmt.Sprintf(":%d", serverPort)
	)

	logger.Setup()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	srv := server.New(httpAddr)

	go func() {
		logger.Infof("Server started at port %d", serverPort)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logger.Fatal(err.Error())
			}
		}
	}()

	// Gracefully shutdown
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown")
	}

	logger.Info("Server exiting")

}
