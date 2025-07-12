package controllers

import (
	"cine-resenha-go/src/dtos"
	"cine-resenha-go/src/repositories"
	"cine-resenha-go/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(server *gin.Engine, repo *repositories.UserRepository) {
	service := services.NewUserService(repo)
	controller := &UserController{service: service}

	routes := server.Group("/users")
	{
		routes.POST("", controller.CreateUser)
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userDto dtos.User

	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Dados invalidos",
		})
		return
	}

	err = uc.service.CreateUser(userDto.Email, userDto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao criar usuario.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario criado com sucesso.",
	})
}
