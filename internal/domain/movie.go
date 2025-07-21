package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Title       string             `bson:"title" json:"title"`
	Summary     string             `bson:"summary" json:"summary"`
	Director    string             `bson:"director" json:"director"`
	CoverLink   string             `bson:"cover_link" json:"cover_link"`
	WhoChose    string             `bson:"who_chose" json:"who_chose"`
	ReleaseYear int                `bson:"release_year" json:"release_year"`
	WatchedDate CustomDate         `bson:"watched_date" json:"watched_date"`
	Rating      float32            `bson:"rating" json:"rating"`
	Duration    int                `bson:"duration" json:"duration"`
	Watched     bool               `bson:"watched" json:"watched"`
}

type MovieRepository interface {
	GetAll(ctx context.Context) ([]Movie, error)
	GetById(ctx context.Context, id primitive.ObjectID) (*Movie, error)
	Create(ctx context.Context, movie *Movie) error
	Update(ctx context.Context, movie *Movie) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
