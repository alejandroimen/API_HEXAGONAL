package controllers

import (
	"log"

	"github.com/alejandroimen/API_HEXAGONAL/src/users/application"
	"github.com/gin-gonic/gin"
)

type GetUsersController struct {
	getUsers *application.GetUsers
}

func NewUsersController(getUsers *application.GetUsers) *GetUsersController {
	return &GetUsersController{getUsers: getUsers}
}

func (gu *GetUsersController) Handle(ctx *gin.Context) {
	log.Println("Petición de listar todos los usuarios, recibido")

	user, err := gu.getUsers.Run()
	if err != nil {
		log.Printf("Error buscando usuarios")
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Retornando|arios", len(user))
	ctx.JSON(200, user)

}
