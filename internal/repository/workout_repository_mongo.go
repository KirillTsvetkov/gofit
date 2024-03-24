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

type WorkoutRepositoryMongo struct {
	db *mongo.Collection
}

func NewWorkoutRepositoryMongo(dbClient *mongo.Database, collectionName string) *WorkoutRepositoryMongo {
	return &WorkoutRepositoryMongo{
		db: dbClient.Collection(collectionName),
	}
}

func (rep *WorkoutRepositoryMongo) CreateWorkout(ctx context.Context, workout *domain.Workout, user *domain.User) (*domain.Workout, error) {
	res, err := rep.db.InsertOne(ctx, workout)
	if err != nil {
		return workout, err
	}
	workout.ID = res.InsertedID.(primitive.ObjectID)
	return workout, nil
}

func (rep *WorkoutRepositoryMongo) GetWorkoutById(ctx context.Context, id string) (*domain.Workout, error) {
	workout := new(domain.Workout)
	err := rep.db.FindOne(ctx, bson.M{
		"ID": id,
	}).Decode(workout)

	if err != nil {
		return nil, err
	}
	return workout, nil
}

func (rep *WorkoutRepositoryMongo) GetWorkoutByUserId(ctx context.Context, user *domain.User) ([]domain.Workout, error) {
	var workout []domain.Workout
	err := rep.db.FindOne(ctx, bson.M{
		"user_id": user.ID,
	}).Decode(workout)

	if err != nil {
		return nil, err
	}
	return workout, nil
}

func (rep *WorkoutRepositoryMongo) UpdateWorkout(ctx context.Context, workout domain.Workout) (*domain.Workout, error) {
	filter := bson.M{"_id": workout.ID}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	findUpdateOptions := options.FindOneAndUpdateOptions{}
	findUpdateOptions.SetReturnDocument(options.After)

	var updatedWorkout domain.Workout
	err := rep.db.FindOneAndUpdate(ctx, filter, update, &findUpdateOptions).Decode(&updatedWorkout)
	if err != nil {
		return nil, err
	}

	return &updatedWorkout, nil
}

func (rep *WorkoutRepositoryMongo) DeleteWorkout(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	_, err = rep.db.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}

func (rep *WorkoutRepositoryMongo) ListWorkouts(ctx context.Context, user *domain.User, query domain.GetWorkoutListQuery) ([]domain.Workout, int64, error) {
	filter := bson.M{"$and": []bson.M{{"user_id": user.ID}}}

	if err := filterDateQueries(query.FilterQuery.DateFrom, query.FilterQuery.DateTo, "date", filter); err != nil {
		return nil, 0, err
	}

	totalCount, err := rep.db.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fOpt := query.PaginationQuery.GetPaginationOpts()

	cursor, err := rep.db.Find(ctx, filter, fOpt)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var workouts []domain.Workout
	for cursor.Next(ctx) {
		var workout domain.Workout
		if err := cursor.Decode(&workout); err != nil {
			log.Println(err)
		}

		workouts = append(workouts, workout)
	}

	return workouts, totalCount, nil
}
