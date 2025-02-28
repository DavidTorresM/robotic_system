package controllers

import (
	"net/http"

	"robotica_concursos/models"
	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type ParticipanteController struct {
	service services.ParticipanteServiceInterface
}

func NewParticipanteController(service services.ParticipanteServiceInterface) *ParticipanteController {
	return &ParticipanteController{service: service}
}

func (pc ParticipanteController) GetParticipantes(c *gin.Context) {
	participantes, err := pc.service.GetParticipantes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, participantes)
}

func (pc ParticipanteController) GetParticipante(c *gin.Context) {
	id := c.Param("id")
	participante, err := pc.service.GetParticipante(id)
	if err != nil {
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
	createdParticipante, err := pc.service.CreateParticipante(participante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdParticipante)
}

func (pc ParticipanteController) UpdateParticipante(c *gin.Context) {
	id := c.Param("id")
	var participante models.Participante
	if _, err := pc.service.GetParticipante(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participante not found"})
		return
	}
	if err := c.ShouldBindJSON(&participante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedParticipante, err := pc.service.UpdateParticipante(participante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedParticipante)
}

func (pc ParticipanteController) DeleteParticipante(c *gin.Context) {
	id := c.Param("id")
	if err := pc.service.DeleteParticipante(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func RegisterParticipanteRoutes(router *gin.Engine, service services.ParticipanteServiceInterface) {
	pc := NewParticipanteController(service)
	router.GET("/participantes", pc.GetParticipantes)
	router.GET("/participantes/:id", pc.GetParticipante)
	router.POST("/participantes", pc.CreateParticipante)
	router.PUT("/participantes/:id", pc.UpdateParticipante)
	router.DELETE("/participantes/:id", pc.DeleteParticipante)
}
