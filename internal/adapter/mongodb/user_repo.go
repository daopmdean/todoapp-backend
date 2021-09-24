package mongodb

import (
	"context"
	"time"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// check if missing implementation
var _ repo.UserRepo = new(userRepo)

func NewUserRepo(col *mongo.Collection) repo.UserRepo {
	return &userRepo{col}
}

type userRepo struct {
	usersCollection *mongo.Collection
}

func (ur *userRepo) CreateUser(user *entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := ur.usersCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) GetUserByUsername(username string) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user entity.User
	err := ur.usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil
	}
	return &user
}
