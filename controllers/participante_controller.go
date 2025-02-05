package controllers

import (
	"net/http"

	"robotica_concursos/models"
	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type ParticipanteController struct{}

func (pc ParticipanteController) GetParticipantes(c *gin.Context) {
	var participantes []models.Participante
	if err := services.GetDatabase().Find(&participantes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, participantes)
}

func (pc ParticipanteController) GetParticipante(c *gin.Context) {
	id := c.Param("id")
	var participante models.Participante
	if err := services.GetDatabase().First(&participante, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participante not found"})
		return
	}
	c.JSON(http.StatusOK, participante)
}

func (pc ParticipanteController) CreateParticipante(c *gin.Context) {
	var participante models.Participante
	if err := c.ShouldBindJSON(&participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.GetDatabase().Create(&participante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, participante)
}

func (pc ParticipanteController) UpdateParticipante(c *gin.Context) {
	id := c.Param("id")
	var participante models.Participante
	if err := services.GetDatabase().First(&participante, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participante not found"})
		return
	}
	if err := c.ShouldBindJSON(&participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.GetDatabase().Save(&participante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, participante)
}

func (pc ParticipanteController) DeleteParticipante(c *gin.Context) {
	id := c.Param("id")
	if err := services.GetDatabase().Delete(&models.Participante{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func RegisterParticipanteRoutes(router *gin.Engine) {
	pc := ParticipanteController{}
	router.GET("/participantes", pc.GetParticipantes)
	router.GET("/participantes/:id", pc.GetParticipante)
	router.POST("/participantes", pc.CreateParticipante)
	router.PUT("/participantes/:id", pc.UpdateParticipante)
	router.DELETE("/participantes/:id", pc.DeleteParticipante)
}
