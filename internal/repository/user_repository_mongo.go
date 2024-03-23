package repository

import (
	"context"
	"log"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryMongo struct {
	db *mongo.Collection
}

func NewUserRepositoryMongo(dbClient *mongo.Database, collectionName string) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		db: dbClient.Collection(collectionName),
	}
}

func (rep *UserRepositoryMongo) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return &user, nil
}

func (rep *UserRepositoryMongo) GetUser(ctx context.Context, email, password string) (*domain.User, error) {
	var user domain.User
	log.Print("Username: " + email + " password: " + password)
	err := rep.db.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (rep *UserRepositoryMongo) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	filter := bson.M{
		"_id": objectId,
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	result := rep.db.FindOne(ctx, filter)

	if err := result.Decode(&user); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func (rep *UserRepositoryMongo) UpdateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if user.Name != "" {
		update["$set"].(bson.M)["name"] = user.Name
	}
	if user.Email != "" {
		update["$set"].(bson.M)["email"] = user.Email
	}

	findUpdateOptions := options.FindOneAndUpdateOptions{}
	findUpdateOptions.SetReturnDocument(options.After)

	var updatedUser domain.User
	err := rep.db.FindOneAndUpdate(ctx, filter, update, &findUpdateOptions).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (rep *UserRepositoryMongo) DeleteUser(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}

	_, err = rep.db.UpdateOne(ctx, filter, update)

	return err
}

func (rep *UserRepositoryMongo) ListUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	return users, nil
}
