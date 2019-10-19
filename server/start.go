package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Start starts the API's webserver
func Start(router *mux.Router) error {

	// server
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// logging
	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	if logFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    500, // in MB
			MaxBackups: 3,
			MaxAge:     28,   // in days
			Compress:   true, // false by default
		})
	}

	// start server
	go func() {
		log.Println("Starting server")
		err := srv.ListenAndServe()
		// https://github.com/denji/golang-tls
		// err := http.ListenAndServeTLS(":8080", "https-server.crt", "https-server.key", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	waitForShutdown(srv)

	return nil
}
