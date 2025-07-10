package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Title       string             `bson:"title" json:"title"`
	Summary     string             `bson:"summary" json:"summary"`
	Director    string             `bson:"director" json:"director"`
	CoverLink   string             `bson:"coverLink" json:"coverLink"`
	WhoChose    string             `bson:"whoChose" json:"whoChose"`
	ReleaseYear int                `bson:"releaseYear" json:"releaseYear"`
	WatchedDate time.Time          `bson:"watchedDate" json:"watchedDate"`
	Rating      float32            `bson:"rating" json:"rating"`
	Duration    int                `bson:"duration" json:"duration"`
	Watched     bool               `bson:"watched" json:"watched"`
}
