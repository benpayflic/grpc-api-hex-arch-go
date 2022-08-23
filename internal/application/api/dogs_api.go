package api

import "github.com/benpayflic/grpc-api-hex-arch-go/internal/application/domain/dogs"

func (api Application) CreateDog(d *dogs.Dog) error {

	api.dogService.CalculateHumanYears(d)

	err := api.db.CreateDog(d)
	if err != nil {
		return err
	}

	return nil
}

func (api Application) RetrieveDog(dogName string) (*dogs.Dog, error) {

	dog, err := api.db.RetrieveDog(dogName)
	if err != nil {
		return nil, err
	}
	return dog, nil
}

func (api Application) UpdateDog(d *dogs.Dog) error {

	api.dogService.CalculateHumanYears(d)

	err := api.db.UpdateDog(d)
	if err != nil {
		return err
	}
	return nil
}

func (api Application) DeleteDog(dogName string) error {

	err := api.db.DeleteDog(dogName)
	if err != nil {
		return err
	}
	return nil
}
