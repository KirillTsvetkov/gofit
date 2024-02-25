package repository

import (
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

func (rep *UserRepositoryMongo) CreateUser(user models.User) (models.User, error) {
	return user, nil
}

func (rep *UserRepositoryMongo) GetUserByID(id string) (models.User, error) {
	var user models.User
	return user, nil
}

func (rep *UserRepositoryMongo) UpdateUser(user models.User) (models.User, error) {
	return user, nil
}

func (rep *UserRepositoryMongo) DeleteUser(id string) error {
	return nil
}

func (rep *UserRepositoryMongo) ListUsers() ([]models.User, error) {
	var users []models.User
	return users, nil
}
