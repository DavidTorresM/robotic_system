package services

import (
	"net/http"
	"robotica_concursos/models"

	"github.com/gin-gonic/gin"
)

func GetEquipos(c *gin.Context, page int, pageSize int) {
	var equipos []models.Equipo
	// Logic to get all equipos
	var db = GetDatabase()
	db.Limit(pageSize).Find(&equipos)
	c.JSON(http.StatusOK, equipos)
}

func GetEquipoByID(c *gin.Context) {
	id := c.Param("id")
	var equipo models.Equipo
	// Logic to get equipo by ID
	var db = GetDatabase()
	db.First(&equipo, id)
	c.JSON(http.StatusOK, equipo)
}

func PostEquipo(c *gin.Context) {
	var equipo models.Equipo
	if err := c.ShouldBindJSON(&equipo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to create a new equipo
	var db = GetDatabase()
	if err := db.Create(&equipo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, equipo)
}

func UpdateEquipo(c *gin.Context) {
	id := c.Param("id")
	var equipo models.Equipo
	if err := c.ShouldBindJSON(&equipo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to update equipo by ID
	var db = GetDatabase()
	db.First(&equipo, id)

	c.JSON(http.StatusOK, equipo)
}

func DeleteEquipo(c *gin.Context) {
	id := c.Param("id")
	// Logic to delete equipo by ID
	var db = GetDatabase()
	db.Delete(&models.Equipo{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Equipo deleted"})
}
