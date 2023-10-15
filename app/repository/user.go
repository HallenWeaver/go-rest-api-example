package repository

import (
	"alexandre/gorest/app/model"
	"context"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	userRepoDatabase := client.Database("charDB")
	userCollection := userRepoDatabase.Collection("users")
	return &UserRepository{userCollection: userCollection}
}

func (ur *UserRepository) CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (bool, error) {
	// Check whether the email provided is valid
	_, err := mail.ParseAddress(newUser.Email)
	if err != nil {
		return false, err
	}

	// Encrypt Password For Proper Storage
	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return false, err
	}
	newUser.Password = string(bytes)

	// Set new ID and a role for User
	newUser.ID = primitive.NewObjectID()
	newUser.Role = role

	// Insert user data
	_, err = ur.userCollection.InsertOne(ctx, newUser)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ur *UserRepository) LoginUser(ctx context.Context, loginUser model.User) (bool, error) {
	user := &model.User{}
	err := ur.userCollection.FindOne(ctx, model.User{Username: loginUser.Username}).Decode(&user)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		return false, err
	}
	return true, nil
}
