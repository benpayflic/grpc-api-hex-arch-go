package ports

import "github.com/benpayflic/grpc-api-hex-arch-go/internal/application/domain/dogs"

// Driving ports

type APIPort interface {
	CreateDog(*dogs.Dog) error
	RetrieveDog(string) (*dogs.Dog, error)
	UpdateDog(*dogs.Dog) error
	DeleteDog(string) error
}
