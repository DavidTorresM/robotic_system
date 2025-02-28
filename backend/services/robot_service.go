package services

import (
	"net/http"
	"robotica_concursos/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RobotService interface {
	GetRobots(c *gin.Context, page int, pageSize int)
	GetRobotByID(c *gin.Context)
	PostRobot(c *gin.Context)
	UpdateRobot(c *gin.Context)
	DeleteRobot(c *gin.Context)
}

type robotService struct {
	db *gorm.DB
}

func NewRobotService(db *gorm.DB) RobotService {
	return &robotService{db: db}
}

func (s *robotService) GetRobots(c *gin.Context, page int, pageSize int) {
	var robots []models.Robot
	// Logic to get all robots
	db := s.db
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&robots)
	for i := 0; i < len(robots); i++ {
		var equipo models.Equipo
		db.Select([]string{}).First(&equipo, "id = ?", robots[i].EquipoID)
		robots[i].Equipo = &equipo
	}
	c.JSON(http.StatusOK, robots)
}

func (s *robotService) GetRobotByID(c *gin.Context) {
	id := c.Param("id")
	var robot models.Robot
	// Logic to get robot by ID
	db := s.db
	result := db.First(&robot, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}

	var equipo models.Equipo
	db.Select([]string{"ID", "Nombre"}).First(&equipo, "id = ?", robot.EquipoID)
	robot.Equipo = &equipo
	c.JSON(http.StatusOK, robot)
}

func (s *robotService) PostRobot(c *gin.Context) {
	var robot models.Robot
	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to create a new robot
	db := s.db
	if err := db.Create(&robot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, robot)
}

func (s *robotService) UpdateRobot(c *gin.Context) {
	id := c.Param("id")
	var robot models.Robot
	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to update robot by ID
	db := s.db
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

func (s *robotService) DeleteRobot(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to delete robot by ID
	db := s.db
	var robot models.Robot
	robot.ID = uint(id)
	db.Delete(&robot)
	c.JSON(http.StatusOK, gin.H{"message": "Robot deleted"})
}
