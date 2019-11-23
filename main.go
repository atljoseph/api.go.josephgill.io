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

func main() {

	// TODO: DON'T PANIC
	// TODO: A real logging solution logrus?

	// flag vars
	isProd := flag.Bool("isProd", false, "set this flag when building prod")
	flag.Parse()

	// init the logger
	loggerConfig := &logger.Config{
		Filename: os.Getenv("LOG_FILENAME")}
	err = logger.Initialize(loggerConfig)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}

	// init the aws connectors
	// singleton package
	awsConfig := &aws.Config{
		S3PublicName:   os.Getenv("S3_PUBLIC_NAME"),
		S3PublicURL:    os.Getenv("S3_PUBLIC_URL"),
		S3PublicSecret: os.Getenv("S3_PUBLIC_SECRET"),
	}
	err = aws.Initialize(awsConfig)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	aws.S3PublicAssetURL("123")

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
		log.Fatal(err)
		// panic(err)
	}

	// TODO: Write authDB and migration

	// configure the routes
	routesConfig := &routes.Config{
		IsProd: *isProd}
	router, err := routes.Initialize(routesConfig)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}

	// start the go server
	err = server.Start(router)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
}
