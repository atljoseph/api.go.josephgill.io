package main

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/routes"
	"github.com/atljoseph/api.josephgill.io/server"
	"github.com/atljoseph/api.josephgill.io/sqlite"
)

func main() {

	// setup the database connection(s) - this is a singleton
	_, err := sqlite.Connect("./i-love-nhung.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

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
