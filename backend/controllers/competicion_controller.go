package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

	// Placeholder for fetching the competition by ID
	// competition := fetchCompetitionByID(id)

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func RegisterRoutesCompeticion(router *gin.Engine) {
	router.GET("/competicion", GetCompeticion)
	router.POST("/competicion/start", StartCompeticion)
}
