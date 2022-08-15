package ports

import (
	d "github.com/benpayflic/grpc-api-hex-arch-go/internal/application/domain/dogs"
)

// Driven ports

type DBPort interface {
	Connect()
	CreateDog(*d.Dog) error
	RetrieveDog(string) (*d.Dog, error)
	UpdateDog(*d.Dog) error
	DeleteDog(string) error
}
