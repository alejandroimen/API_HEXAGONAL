package controllers

import (
	"strconv"

	"github.com/alejandroimen/API_HEXAGONAL/users/application"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUser *application.UpdateUser
}

func NewUpdateUserController(updateUser *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{updateUser: updateUser}
}

func (update *UpdateUserController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de usuario invalido"})
		return
	}

	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "petición del body invalida"})
		return
	}

	if err := update.updateUser.Run(id, request.Email, request.Name, request.Password); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "usuario actualizado correctamente"})
}
