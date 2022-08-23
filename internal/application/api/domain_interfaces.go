package api

import "github.com/benpayflic/grpc-api-hex-arch-go/internal/application/domain/dogs"

type DogService interface {
	CalculateHumanYears(d *dogs.Dog)
}
