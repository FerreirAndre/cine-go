package repository

import (
	"context"

	"github.com/ferreirandre/cine-go/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type movieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(db *mongo.Database) domain.MovieRepository {
	return &movieRepository{
		collection: db.Collection("movies"),
	}
}

// Get implements domain.MovieRepository
func (r *movieRepository) GetAll(ctx context.Context) ([]domain.Movie, error) {
	var movies []domain.Movie
	filter := bson.M{}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return movies, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var movie domain.Movie
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

// Create implements domain.MovieRepository.
func (r *movieRepository) Create(ctx context.Context, movie *domain.Movie) error {
	movie.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, movie)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.MovieRepository.
func (r *movieRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

// GetById implements domain.MovieRepository.
func (r *movieRepository) GetById(ctx context.Context, id primitive.ObjectID) (*domain.Movie, error) {
	var movie domain.Movie
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&movie)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &movie, err
}

// Update implements domain.MovieRepository.
func (r *movieRepository) Update(ctx context.Context, movie *domain.Movie) error {
	filter := bson.M{"_id": movie.ID}
	update := bson.M{
		"$set": bson.M{
			"title":        movie.Title,
			"summary":      movie.Summary,
			"director":     movie.Director,
			"cover_link":   movie.CoverLink,
			"who_chose":    movie.WhoChose,
			"release_year": movie.ReleaseYear,
			"watched_date": movie.WatchedDate,
			"rating":       movie.Rating,
			"duration":     movie.Duration,
			"watched":      movie.Watched,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// ToggleWatched implements domain.MovieRepository.
func (r *movieRepository) ToggleWatched(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$bit": bson.M{
			"watched": bson.M{"xor": 1},
		},
	}

	result := r.collection.FindOneAndUpdate(ctx, filter, update)

	return result.Err()
}
