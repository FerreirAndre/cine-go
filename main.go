package main

import (
	"cine-resenha-go/src/controllers"
	"cine-resenha-go/src/repositories"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	uri, dbName := os.Getenv("URI"), os.Getenv("DBNAME")

	userRepository, err := repositories.NewUserRepository(uri, dbName, "user")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = repositories.NewMovieRepository(uri, dbName, "movies")
	if err != nil {
		log.Fatal(err)
		return
	}

	server := gin.Default()

	controllers.NewUserController(server, userRepository)

	server.Run(":8080")
}
