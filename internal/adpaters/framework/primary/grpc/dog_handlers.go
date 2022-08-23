package grpc

import (
	"context"
	"fmt"

	pb "github.com/benpayflic/grpc-api-hex-arch-go/internal/adpaters/framework/primary/grpc/proto"
	d "github.com/benpayflic/grpc-api-hex-arch-go/internal/application/domain/dogs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDog(ctx context.Context, in *pb.CreateDogRequest) (*pb.CreateDogResponse, error) {

	dog := d.NewDog()
	dog.Name = in.Name
	dog.Breed = in.Breed
	dog.DOB = in.Dob

	err := s.adapter.api.CreateDog(dog)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	return &pb.CreateDogResponse{
		Message: fmt.Sprintf("Created new dog with name %s", in.Name),
	}, nil
}

func (s *Server) RetrieveDog(ctx context.Context, in *pb.RetrieveDogRequest) (*pb.RetrieveDogResponse, error) {

	dog, err := s.adapter.api.RetrieveDog(in.Name)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can not find dog with provided name: %v", in.Name),
		)
	}

	return &pb.RetrieveDogResponse{
		Message: "Found dog!",
		Dog: &pb.RetrieveDogResponse_Dog{
			Name:       dog.Name,
			Breed:      dog.Breed,
			Dob:        dog.DOB,
			HumanYears: dog.HumanYears,
		},
	}, nil

}

func (s *Server) UpdateDog(ctx context.Context, in *pb.UpdateDogRequest) (*pb.UpdateDogResponse, error) {

	dog := d.NewDog()
	dog.Name = in.Name
	dog.Breed = in.Breed
	dog.DOB = in.Dob

	err := s.adapter.api.UpdateDog(dog)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	return &pb.UpdateDogResponse{
		Message: fmt.Sprintf("Updated new dog with name %s", in.Name),
	}, nil
}

func (s *Server) DeleteDog(ctx context.Context, in *pb.DeleteDogRequest) (*pb.DeleteDogResponse, error) {

	err := s.adapter.api.DeleteDog(in.Name)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	return &pb.DeleteDogResponse{
		Message: "Dog deleted!",
	}, nil
}
