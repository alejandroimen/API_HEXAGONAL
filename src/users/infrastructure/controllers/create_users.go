package controllers

import (
	"log"

	"github.com/alejandroimen/API_HEXAGONAL/users/application"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	CreateUsers *application.CreateUsers
}

func NewCreateUserController(CreateUsers *application.CreateUsers) *CreateUserController {
	return &CreateUserController{CreateUsers: CreateUsers}
}

func (c *CreateUserController) Handle(ctx *gin.Context) {
	log.Println("Petición de crear un producto, recibido")

	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		log.Printf("Error decodificando la petición del body: %v", err)
		ctx.JSON(400, gin.H{"error": "petición del body invlida"})
		return
	}
	log.Printf("Creando usuario: Name=%s, email=%s", request.Name, request.Email)

	if err := c.CreateUsers.Run(request.Email, request.Name, request.Password); err != nil {
		log.Printf("Error creando el usuario: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Usuario creado exitosamente")
	ctx.JSON(201, gin.H{"message": "usuario creado exitosamente"})

}
