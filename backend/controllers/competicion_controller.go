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

func StartCompeticion(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Placeholder for starting the competition by ID
	err = services.StartCompetitionByID(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Competition started", "id": id})
}

func GetCompeticion(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if id == 1 {
		ronda, err := services.GetRondaCompeticionSumo()
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
		ronda, err := services.GetRondaCompeticionSigueLineas()
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

func FijaGanadorCompeticionSumo(c *gin.Context) {
	var requestBody vo.RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := services.SetWinnerSumo(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Se fijo el ganador de la competicion satisfactoriamente"})
}

func RegisterRoutesCompeticion(router *gin.Engine) {
	router.POST("/competicion/sumo/ganador", FijaGanadorCompeticionSumo)
	router.GET("/competicion", GetCompeticion)
	router.POST("/competicion/start", StartCompeticion)
}
