package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"robotica_concursos/controllers/vo"
	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

// CompeticionController handles competition-related requests
type CompeticionController struct {
	service services.CompeticionService
}

// NewCompeticionController creates a new CompeticionController
func NewCompeticionController(service services.CompeticionService) *CompeticionController {
	return &CompeticionController{service: service}
}

func (cc *CompeticionController) StartCompeticion(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Start the competition by ID
	err = cc.service.StartCompetitionByID(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Competition started", "id": id})
}

func (cc *CompeticionController) GetCompeticion(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if id == 1 {
		ronda, err := cc.service.GetRondaCompeticionSumo()
		if err != nil {
			if strings.Contains(err.Error(), "no hay mas rondas") {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint(err)})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"ronda": *ronda})
		return
	} else if id == 3 {
		ronda, err := cc.service.GetRondaCompeticionSigueLineas()
		if err != nil {
			if strings.Contains(err.Error(), "no hay mas rondas") {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint(err)})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"ronda": *ronda})
		return
	}
}

func (cc *CompeticionController) FijaGanadorCompeticionSumo(c *gin.Context) {
	var requestBody vo.RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := cc.service.SetWinnerSumo(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Se fijo el ganador de la competicion satisfactoriamente"})
}

func RegisterRoutesCompeticion(router *gin.Engine, service services.CompeticionService) {
	controller := NewCompeticionController(service)
	router.POST("/competicion/sumo/ganador", controller.FijaGanadorCompeticionSumo)
	router.GET("/competicion", controller.GetCompeticion)
	router.POST("/competicion/start", controller.StartCompeticion)
}
