package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Title       string             `bson:"title" json:"title" validate:"required,min=2,max=200"`
	Summary     string             `bson:"summary" json:"summary" validate:"required"`
	Director    string             `bson:"director" json:"director" validate:"required"`
	CoverLink   string             `bson:"cover_link" json:"cover_link" validate:"url"`
	WhoChose    string             `bson:"who_chose" json:"who_chose" validate:"required"`
	ReleaseYear int                `bson:"release_year" json:"release_year" validate:"gte=1887,lte=2100"`
	WatchedDate CustomDate         `bson:"watched_date" json:"watched_date"`
	Rating      float32            `bson:"rating" json:"rating" validate:"gte=0,lte=10"`
	Duration    int                `bson:"duration" json:"duration" validate:"gt=0"`
	Watched     bool               `bson:"watched" json:"watched"`
}

type MovieRepository interface {
	GetAll(ctx context.Context) ([]Movie, error)
	GetById(ctx context.Context, id primitive.ObjectID) (*Movie, error)
	Create(ctx context.Context, movie *Movie) error
	Update(ctx context.Context, movie *Movie) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
