package main

import (
	"log"

	rightDB "github.com/benpayflic/grpc-api-hex-arch-go/internal/adpaters/framework/secondary/database"
	c "github.com/benpayflic/grpc-api-hex-arch-go/pkg/config"
)

func main() {

	// config, err := c.LoadConfig("./pkg/config/")
	config, err := c.LoadConfig("/Users/bengoodenough/Documents/Portfolio-Projects/grpc-api-hex-arch-go/pkg/config/")
	if err != nil {
		log.Fatalln("Failed to load application configurations : ", err)
	}

	dbDrivenAdapter, err := rightDB.NewAdapter(&config)
	if err != nil {
		log.Fatalf("failed to initialize db driven adapter: %v", err)
	}
	dbDrivenAdapter.Connect()

}
