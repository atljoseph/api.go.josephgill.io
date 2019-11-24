package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func waitForShutdown(srv *http.Server) {
	funcTag := "waitForShutdown"

	// create a channel
	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// block until signal is received
	<-sigquit

	// create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	logMessage(funcTag, "Shutting down")
	os.Exit(0)
}
