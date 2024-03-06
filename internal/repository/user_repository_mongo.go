package repository

import (
	"context"
	"log"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	db *mongo.Collection
}

func NewUserRepositoryMongo(dbClient *mongo.Database, collectionName string) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		db: dbClient.Collection(collectionName),
	}
}

func (rep *UserRepositoryMongo) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return &user, nil
}

func (rep *UserRepositoryMongo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
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

func (rep *UserRepositoryMongo) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return &user, nil
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

func (rep *UserRepositoryMongo) ListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	return users, nil
}
