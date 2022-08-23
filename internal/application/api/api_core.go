package api

import (
	"github.com/benpayflic/grpc-api-hex-arch-go/internal/ports"
	"github.com/benpayflic/grpc-api-hex-arch-go/pkg/config"
)

// Application implements the APIPort interface
type Application struct {
	db         ports.DBPort
	dogService DogService
	config     config.Config
}

// Return new instance of Application
func NewApplication(db ports.DBPort, dogService DogService, config config.Config) *Application {
	return &Application{
		db:         db,
		dogService: dogService,
		config:     config,
	}
}
