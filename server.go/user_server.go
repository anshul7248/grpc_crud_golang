package server

import (
	"context"
	"errors"
	"grpc_mongo_crud/db"
	_ "grpc_mongo_crud/db"
	pb "grpc_mongo_crud/grpc_mongo_crud/proto"
	"grpc_mongo_crud/models"
	_ "grpc_mongo_crud/models"

	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := models.User{
		UserId:   req.UserId,
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := db.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := models.User{}
	err := db.UserCollection.FindOne(ctx, bson.M{"user_id": req.UserId}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &pb.UserResponse{
		UserId: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{
		"name":     req.Name,
		"age":      req.Age,
		"email":    req.Email,
		"Password": req.Password,
	}}
	_, err := db.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil
}
