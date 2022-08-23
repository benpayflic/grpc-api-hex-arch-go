package grpc

import (
	"log"
	"net"

	pb "github.com/benpayflic/grpc-api-hex-arch-go/internal/adpaters/framework/primary/grpc/proto"
	"github.com/benpayflic/grpc-api-hex-arch-go/internal/ports"
	"github.com/benpayflic/grpc-api-hex-arch-go/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr = "0.0.0.0:8080"

type Adapter struct {
	api    ports.APIPort
	config *config.Config
}

func NewAdapter(api ports.APIPort, config *config.Config) *Adapter {
	return &Adapter{api: api, config: config}
}

type Server struct {
	pb.DogServiceServer

	adapter Adapter
}

func (adapter Adapter) Start() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen: ", addr)
	}

	log.Println("Listening on ", addr)

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterDogServiceServer(s, &Server{adapter: adapter})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve", err)
	}

}
