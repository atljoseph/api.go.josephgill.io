package main

import (
	"github.com/atljoseph/api.josephgill.io/photoDB"
	"github.com/atljoseph/api.josephgill.io/routes"
	"github.com/atljoseph/api.josephgill.io/server"
)

func main() {

	// TODO: DON'T PANIC

	// setup the database connection(s) keychain, as a singleton
	dbConfig := &photoDB.Config{
		MaxOpenConns:    15,
		Username:        "root",
		Password:        "password",
		Host:            "localhost",
		Port:            3306,
		DefaultDatabase: "photos",
	}
	err := photoDB.Initialize(dbConfig)
	if err != nil {
		panic(err)
	}

	// TODO: Write authDB and migration

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
