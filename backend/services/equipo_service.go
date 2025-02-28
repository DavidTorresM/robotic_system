package services

import (
	"net/http"
	"robotica_concursos/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EquipoService defines the interface for equipo service
type EquipoService interface {
	GetEquipos(c *gin.Context, page, pageSize int)
	GetEquipoByID(c *gin.Context)
	PostEquipo(c *gin.Context)
	UpdateEquipo(c *gin.Context)
	DeleteEquipo(c *gin.Context)
}

type equipoService struct {
	db *gorm.DB
}

// NewEquipoService creates a new instance of EquipoService
func NewEquipoService(db *gorm.DB) EquipoService {
	return &equipoService{db: db}
}

func (s *equipoService) GetEquipos(c *gin.Context, page, pageSize int) {
	var equipos []models.Equipo
	s.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&equipos)

	for i := range equipos {
		loadEquipoRelations(&equipos[i], s.db)
	}

	c.JSON(http.StatusOK, equipos)
}

func (s *equipoService) GetEquipoByID(c *gin.Context) {
	id := c.Param("id")
	var equipo models.Equipo

	if err := s.db.First(&equipo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}

	loadEquipoRelations(&equipo, s.db)
	c.JSON(http.StatusOK, equipo)
}

func (s *equipoService) PostEquipo(c *gin.Context) {
	var equipo models.Equipo
	if err := c.ShouldBindJSON(&equipo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Create(&equipo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, equipo)
}

func (s *equipoService) UpdateEquipo(c *gin.Context) {
	id := c.Param("id")
	var equipo models.Equipo
	if err := c.ShouldBindJSON(&equipo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.First(&models.Equipo{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}

	if err := s.db.Model(&equipo).Updates(&equipo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipo)
}

func (s *equipoService) DeleteEquipo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var equipo models.Equipo
	equipo.ID = uint(id)

	s.db.Find(&equipo.Participantes, "equipo_id = ?", id)
	s.db.Find(&equipo.Robots, "equipo_id = ?", id)
	s.db.Delete(&equipo.Participantes)
	s.db.Delete(&equipo.Robots)
	s.db.Delete(&equipo)

	c.JSON(http.StatusOK, gin.H{"message": "Equipo deleted"})
}

func loadEquipoRelations(equipo *models.Equipo, db *gorm.DB) {
	db.Select([]string{"ID", "Nombre", "Correo", "Telefono"}).Find(&equipo.Participantes, "equipo_id = ?", equipo.ID)
	db.Select([]string{"ID", "Nombre", "Descripcion", "EquipoID", "CategoriaID"}).Find(&equipo.Robots, "equipo_id = ?", equipo.ID)
}
