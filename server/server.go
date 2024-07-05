package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-user-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []pb.User
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.GetUserResponse{User: &user}, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s *server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var users []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				users = append(users, &user)
			}
		}
	}
	return &pb.GetUsersResponse{Users: users}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var users []*pb.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Married == user.Married) {
			users = append(users, &user)
		}
	}
	return &pb.SearchUsersResponse{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 1234567891, Height: 5.9, Married: false},
			// Add more users as needed
		},
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
