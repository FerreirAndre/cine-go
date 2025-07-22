package main

import (
	"os"

	"github.com/ferreirandre/cine-go/internal/db"
	"github.com/ferreirandre/cine-go/internal/handler"
	"github.com/ferreirandre/cine-go/internal/repository"
	"github.com/ferreirandre/cine-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	uri, dbName := os.Getenv("URI"), os.Getenv("DBNAME")

	movieCollection, err := db.ConnectMongoDB(uri, dbName)
	if err != nil {
		panic(err)
	}
	repository := repository.NewMovieRepository(movieCollection)

	service := service.NewMovieService(repository)

	movieHandler := handler.NewMovieHandler(service)

	r := gin.Default()

	movies := r.Group("/movies")
	{
		movies.GET("", movieHandler.GetAll)
		movies.GET("/:id", movieHandler.GetById)
		movies.POST("", movieHandler.Create)
		movies.PUT("/:id", movieHandler.Update)
		movies.DELETE("/:id", movieHandler.Delete)
		movies.PATCH("/:id/watched", movieHandler.ToggleWatched)
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
