package handler

import (
	"net/http"

	"github.com/ferreirandre/cine-go/internal/domain"
	"github.com/ferreirandre/cine-go/internal/service"
	"github.com/ferreirandre/cine-go/internal/validation"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type movieHandler struct {
	service service.MovieService
}

func NewMovieHandler(service service.MovieService) *movieHandler {
	return &movieHandler{service: service}
}

func (h *movieHandler) GetAll(c *gin.Context) {
	movies, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (h *movieHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *movieHandler) Create(c *gin.Context) {
	var movie domain.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validation.ValidateMovie(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	err := h.service.Create(c.Request.Context(), &movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *movieHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var movie domain.Movie

	err = c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validation.ValidateMovie(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}

	movie.ID = id
	err = h.service.Update(c.Request.Context(), &movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *movieHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *movieHandler) ToggleWatched(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.ToggleWatched(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
