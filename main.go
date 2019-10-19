package main

import (
	"github.com/atljoseph/api.josephgill.io/routes"
	"github.com/atljoseph/api.josephgill.io/server"
)

func main() {

	// configure the routes
	router, err := routes.Configure()
	if err != nil {
		panic(err)
	}

	// start the go server
	err = server.Start(router)
	if err != nil {
		panic(err)
	}
}
