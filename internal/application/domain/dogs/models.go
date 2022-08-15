package dogs

import "encoding/json"

type Dog struct {
	Name       string `json:"name,omitempty"`
	Breed      string `json:"breed,omitempty"`
	DOB        string `json:"dob,omitempty"`
	HumanYears int64  `json:"humanYears,omitempty"`
	DogFact    string `json:"fact,omitempty"`
}

func UnmarshalDog(data []byte) (Dog, error) {
	var r Dog
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Dog) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func NewDog() *Dog {
	return &Dog{}
}
