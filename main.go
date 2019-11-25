package main

import (
	"flag"
	"log"
	"os"

	"github.com/atljoseph/api.josephgill.io/aws"
	"github.com/atljoseph/api.josephgill.io/logger"
	"github.com/atljoseph/api.josephgill.io/photoDB"
	"github.com/atljoseph/api.josephgill.io/routes"
	"github.com/atljoseph/api.josephgill.io/server"
)

var err error
var mainLog *logger.Log

func main() {
	funcTag := "main"

	// TODO: DON'T PANIC

	// flag vars
	isProd := flag.Bool("isProd", false, "set this flag when building prod")
	flag.Parse()

	// init the logger
	// this MUST be done first
	loggerConfig := &logger.Config{} // Filename: os.Getenv("LOG_FILENAME")
	err = logger.Initialize(loggerConfig)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	mainLog = logger.ForPackage("main").WithFunc(funcTag)

	// init the aws connectors
	// singleton package
	awsConfig := &aws.Config{
		S3PublicRegion: os.Getenv("S3_PUBLIC_REGION"),
		S3PublicName:   os.Getenv("S3_PUBLIC_NAME"),
		S3PublicURL:    os.Getenv("S3_PUBLIC_URL"),
		S3UserID:       os.Getenv("S3_USER_ID"),
		S3UserSecret:   os.Getenv("S3_USER_SECRET"),
	}
	err = aws.Initialize(awsConfig)
	if err != nil {
		mainLog.WithError(err).Panic()
	}

	// init the photo db
	// singleton package
	dbConfig := &photoDB.Config{
		MaxOpenConns:    15,
		Username:        os.Getenv("PHOTODB_USER"),
		Password:        os.Getenv("PHOTODB_PASS"),
		Host:            os.Getenv("PHOTODB_HOST"),
		Port:            3306,
		DefaultDatabase: "photos",
	}
	err = photoDB.Initialize(dbConfig)
	if err != nil {
		mainLog.WithError(err).Panic()
	}

	// TODO: Write authDB and migration

	// configure the routes
	routesConfig := &routes.Config{
		IsProd: *isProd}
	router, err := routes.Initialize(routesConfig)
	if err != nil {
		mainLog.WithError(err).Panic()
	}

	// start the go server
	err = server.Start(router)
	if err != nil {
		mainLog.WithError(err).Panic()
	}
}
