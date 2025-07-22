package validation

import (
	"github.com/ferreirandre/cine-go/internal/domain"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateMovie(movie *domain.Movie) error {
	return validate.Struct(movie)
}
