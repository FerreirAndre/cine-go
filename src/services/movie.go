package services

import "cine-resenha-go/src/repositories"

type MovieService struct {
	repo *repositories.MovieRepository
}

func NewMovieService(repo *repositories.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}
