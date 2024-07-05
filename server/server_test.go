package main

import (
	"context"
	"testing"

	pb "grpc-user-service/proto"
)

func TestGetUser(t *testing.T) {
	s := &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		},
	}

	req := &pb.GetUserRequest{Id: 1}
	res, err := s.GetUser(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.User.Id != 1 {
		t.Fatalf("expected user ID 1, got %v", res.User.Id)
	}
}

func TestSearchUsers(t *testing.T) {
	s := &server{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 1234567891, Height: 5.9, Married: false},
		},
	}

	req := &pb.SearchUsersRequest{City: "LA"}
	res, err := s.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res.Users) != 1 || res.Users[0].Id != 1 {
		t.Fatalf("expected one user with ID 1, got %v", res.Users)
	}
}
