package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Start starts the API's webserver
func Start(router *mux.Router) error {
	funcTag := "server.Start"

	// server
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start server
	go func() {
		// log
		pkgLog.WithFunc(funcTag).WithMessage("Starting up").Info()

		// serve
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
