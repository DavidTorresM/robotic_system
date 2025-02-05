package controllers

import (
	"net/http"
	"robotica_concursos/models"
	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

func RegisterParticipante(c *gin.Context) {
	var participante models.Participante
	if err := c.ShouldBindJSON(&participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := services.GetDatabase()
	if err := db.Create(&participante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	participante.Password = ""
	c.JSON(http.StatusOK, gin.H{"data": participante})
}

func RegisterRegistreRoutes(router *gin.Engine) {
	router.POST("/register", RegisterParticipante)
}
