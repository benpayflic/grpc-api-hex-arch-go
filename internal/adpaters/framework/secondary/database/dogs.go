package database

import (
	"context"
	"fmt"

	d "github.com/benpayflic/grpc-api-hex-arch-go/internal/application/domain/dogs"
	"go.mongodb.org/mongo-driver/bson"
)

func (dba *Adapter) CreatDog(dog *d.Dog) error {
	_, err := dba.collection.InsertOne(context.Background(), dog)
	if err != nil {
		return err
	}
	return nil
}

func (dba *Adapter) RetrieveDog(name string) (*d.Dog, error) {
	data := d.NewDog()
	filter := bson.M{"name": name}
	res := dba.collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, fmt.Errorf("no dog found with name %v", name)
	}
	return data, nil
}

func (dba *Adapter) UpdateDog(dog *d.Dog) error {

	filter := bson.M{"namr": dog.Name}

	res, err := dba.collection.UpdateOne(context.Background(),
		filter,
		bson.M{"$set": dog},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("can not find blog with provided name")
	}
	return nil
}

func (dba *Adapter) DeleteDog(name string) error {

	res, err := dba.collection.DeleteOne(context.Background(), bson.M{"name": name})
	if err != nil {
		return fmt.Errorf("can not delete")
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no dog found with provided name")
	}

	return nil
}
