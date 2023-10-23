package repository

import (
	"alexandre/gorest/app/model"
	"context"
	"errors"
	"net/mail"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("mysuperawesomekey")

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	userRepoDatabase := client.Database("charDB")
	userCollection := userRepoDatabase.Collection("users")
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

// GenerateToken generates a jwt token
// GenerateToken takes an email as an argument and returns a signed JWT token and an error
func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &model.JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*model.JWTClaim)
	if !ok {
		err = errors.New("unable to parse JWT claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
