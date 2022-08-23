package dogs

import "github.com/benpayflic/grpc-api-hex-arch-go/pkg/utils"

type DogService struct {
}

func NewDogService() *DogService {
	return &DogService{}
}

// Calculate the age of the dog in human years
func (DogService) CalculateHumanYears(d *Dog) {
	age := utils.Age(d.DOB)
	var humanYears int64

	if age > 6 {
		age = age - 6
		age = age * 4
		humanYears = int64(age * 40)
	} else {
		age = age * 19
		age = age / 3
		humanYears = int64(age + 1)
	}
	d.HumanYears = humanYears
}
