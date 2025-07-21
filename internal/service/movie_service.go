package service

import (
	"context"

	"github.com/ferreirandre/cine-go/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieService interface {
	GetAll(ctx context.Context) ([]domain.Movie, error)
	GetById(ctx context.Context, id primitive.ObjectID) (*domain.Movie, error)
	Create(ctx context.Context, movie *domain.Movie) error
	Update(ctx context.Context, movie *domain.Movie) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

func NewMovieService(repo domain.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

type movieService struct {
	repo domain.MovieRepository
}

// Create implements MovieService.
func (m *movieService) Create(ctx context.Context, movie *domain.Movie) error {
	return m.repo.Create(ctx, movie)
}

// Delete implements MovieService.
func (m *movieService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return m.repo.Delete(ctx, id)
}

// GetAll implements MovieService.
func (m *movieService) GetAll(ctx context.Context) ([]domain.Movie, error) {
	return m.repo.GetAll(ctx)
}

// GetById implements MovieService.
func (m *movieService) GetById(ctx context.Context, id primitive.ObjectID) (*domain.Movie, error) {
	return m.repo.GetById(ctx, id)
}

// Update implements MovieService.
func (m *movieService) Update(ctx context.Context, movie *domain.Movie) error {
	return m.repo.Update(ctx, movie)
}
