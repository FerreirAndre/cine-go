package repositories

type MovieRepository struct {
	*MongoRepositoryContext
}

func NewMovieRepository(uri, dbName, collectionName string) (*MovieRepository, error) {
	mongoRepo, err := NewMongoRepositoryContext(uri, dbName, collectionName)
	if err != nil {
		return nil, err
	}

	return &MovieRepository{MongoRepositoryContext: mongoRepo}, nil
}
