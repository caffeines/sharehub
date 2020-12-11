package data

import (
	"context"
	"time"

	"github.com/caffeines/filepile/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepoImpl ...
type UserRepoImpl struct {
}

var userRepo UserRepository

func NewUserRepo() UserRepository {
	if userRepo == nil {
		userRepo = &UserRepoImpl{}
	}
	return userRepo
}

// Register creates new user
func (usr *UserRepoImpl) Register(db *mongo.Database, user *models.User) (*mongo.InsertOneResult, error) {
	userCollection := db.Collection(user.CollectionName())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	createdUser, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

//FindUserByEmail returns the matched document with email
func (usr *UserRepoImpl) FindUserByEmail(db *mongo.Database, email string) (*models.User, error) {
	user := models.User{}
	userCollection := db.Collection(user.CollectionName())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
