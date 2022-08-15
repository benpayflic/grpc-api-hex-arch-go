package database

import (
	"github.com/benpayflic/grpc-api-hex-arch-go/pkg/config"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adapter struct {
	collection *mongo.Collection
	config     *config.Config
}

func NewAdapter(config *config.Config) (*Adapter, error) {
	collection := &mongo.Collection{}
	return &Adapter{collection: collection, config: config}, nil
}
