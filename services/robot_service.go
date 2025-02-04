package services

import (
	"net/http"
	"robotica_concursos/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRobots(c *gin.Context, page int, pageSize int) {
	var robots []models.Robot
	// Logic to get all robots
	var db = GetDatabase()
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&robots)
	for i := 0; i < len(robots); i++ {
		var Equipo models.Equipo
		db.Select([]string{}).First(&Equipo, "id = ?", robots[i].EquipoID)
		robots[i].Equipo = &Equipo
	}
	c.JSON(http.StatusOK, robots)
}

func GetRobotByID(c *gin.Context) {
	id := c.Param("id")
	var robot models.Robot
	// Logic to get robot by ID
	var db = GetDatabase()
	result := db.First(&robot, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}

	var Equipo models.Equipo
	db.Select([]string{"ID", "Nombre"}).First(&Equipo, "id = ?", robot.EquipoID)
	robot.Equipo = &Equipo
	c.JSON(http.StatusOK, robot)
}

func PostRobot(c *gin.Context) {
	var robot models.Robot
	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to create a new robot
	var db = GetDatabase()
	if err := db.Create(&robot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, robot)
}

func UpdateRobot(c *gin.Context) {
	id := c.Param("id")
	var robot models.Robot
	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to update robot by ID
	var db = GetDatabase()
	result := db.First(&models.Robot{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}
	if err := db.Model(&robot).Updates(&robot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, robot)
}

func DeleteRobot(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to delete robot by ID
	var db = GetDatabase()
	var robot models.Robot
	robot.ID = uint(id)
	db.Delete(&robot)
	c.JSON(http.StatusOK, gin.H{"message": "Robot deleted"})
}
