package services

import (
	"net/http"
	"robotica_concursos/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEquipos(c *gin.Context, page int, pageSize int) {
	var equipos []models.Equipo
	// Logic to get all equipos
	var db = GetDatabase()
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&equipos)
	for i := 0; i < len(equipos); i++ {
		var Participantes []models.Participante
		db.Select([]string{"ID", "Nombre", "Correo", "Telefono"}).Find(&Participantes, "equipo_id = ?", equipos[i].ID)
		var Robots []models.Robot
		db.Select([]string{"ID", "Nombre", "Descripcion", "EquipoID", "CategoriaID"}).Find(&Robots, "equipo_id = ?", equipos[i].ID)
		equipos[i].Participantes = Participantes
		equipos[i].Robots = Robots
	}
	c.JSON(http.StatusOK, equipos)
}

func GetEquipoByID(c *gin.Context) {
	id := c.Param("id")
	var equipo models.Equipo
	// Logic to get equipo by ID
	var db = GetDatabase()
	result := db.First(&equipo, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}

	var Participantes []models.Participante
	db.Select([]string{"ID", "Nombre", "Correo", "Telefono"}).Find(&Participantes, "equipo_id = ?", id)
	var Robots []models.Robot
	db.Select([]string{"ID", "Nombre", "Descripcion", "EquipoID", "CategoriaID"}).Find(&Robots, "equipo_id = ?", id)
	equipo.Participantes = Participantes
	equipo.Robots = Robots
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
	result := db.First(&models.Equipo{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}
	if err := db.Model(&equipo).Updates(&equipo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, equipo)
}

func DeleteEquipo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to delete equipo by ID
	var db = GetDatabase()
	var equipo models.Equipo
	equipo.ID = uint(id)
	db.Find(&equipo.Participantes, "equipo_id = ?", id)
	db.Find(&equipo.Robots, "equipo_id = ?", id)
	db.Delete(&equipo.Participantes)
	db.Delete(&equipo.Robots)
	db.Delete(&equipo)
	c.JSON(http.StatusOK, gin.H{"message": "Equipo deleted"})
}
