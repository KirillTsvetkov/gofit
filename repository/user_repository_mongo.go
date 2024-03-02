package repository

import (
	"context"

	"github.com/KirillTsvetkov/gofit/models"
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
	return &user, nil
}

func (rep *UserRepositoryMongo) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return &user, nil
}

func (rep *UserRepositoryMongo) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (rep *UserRepositoryMongo) ListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	return users, nil
}
