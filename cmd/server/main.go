package main

import (
	"log"

	"github.com/savchenkodn79/go-grpc-services-course/internal/db"
	"github.com/savchenkodn79/go-grpc-services-course/internal/rocket"
)

func Run() error {
	// responsible for initializing and starting
	// out gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migations")
		return err
	}
	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
