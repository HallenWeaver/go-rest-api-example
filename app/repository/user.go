package repository

import (
	"alexandre/gorest/app/model"
	"context"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (*model.User, error)
	LoginUser(ctx context.Context, loginUser model.TokenRequest) (*model.User, error)
}

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	userRepoDatabase := client.Database("charDB")
	userCollection := userRepoDatabase.Collection("users")
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	userCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{emailIndex, usernameIndex})
	return &UserRepository{userCollection: userCollection}
}

func (ur *UserRepository) CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (*model.User, error) {
	// Check whether the email provided is valid
	_, err := mail.ParseAddress(newUser.Email)
	if err != nil {
		return &model.User{}, err
	}

	// Encrypt Password For Proper Storage
	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return &model.User{}, err
	}
	newUser.Password = string(bytes)

	// Set new ID, active status and a role for User
	newUser.ID = primitive.NewObjectID()
	newUser.Role = role
	newUser.IsActive = true

	// Insert user data
	_, err = ur.userCollection.InsertOne(ctx, newUser)

	if err != nil {
		return &model.User{}, err
	}

	return &newUser, nil
}

func (ur *UserRepository) LoginUser(ctx context.Context, loginUser model.TokenRequest) (*model.User, error) {
	// Validate user upon logging in
	user := &model.User{}
	err := ur.userCollection.FindOne(ctx, model.User{Username: loginUser.Username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
