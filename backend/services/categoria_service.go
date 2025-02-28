package services

import (
	"net/http"
	"robotica_concursos/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CategoriaService defines the interface for categoria operations
type CategoriaService interface {
	GetCategorias(page int, pageSize int) ([]models.Categoria, error)
	GetCategoriaByID(id string) (models.Categoria, error)
	CreateCategoria(categoria models.Categoria) error
	UpdateCategoria(id string, categoria models.Categoria) error
	DeleteCategoria(id int) error
}

// categoriaServiceImpl is the implementation of CategoriaService
type categoriaServiceImpl struct {
	db *gorm.DB
}

// NewCategoriaService creates a new CategoriaService
func NewCategoriaService(db *gorm.DB) CategoriaService {
	return &categoriaServiceImpl{db: db}
}

func (s *categoriaServiceImpl) GetCategorias(page int, pageSize int) ([]models.Categoria, error) {
	var categorias []models.Categoria
	err := s.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&categorias).Error
	return categorias, err
}

func (s *categoriaServiceImpl) GetCategoriaByID(id string) (models.Categoria, error) {
	var categoria models.Categoria
	err := s.db.First(&categoria, id).Error
	return categoria, err
}

func (s *categoriaServiceImpl) CreateCategoria(categoria models.Categoria) error {
	return s.db.Create(&categoria).Error
}

func (s *categoriaServiceImpl) UpdateCategoria(id string, categoria models.Categoria) error {
	var existingCategoria models.Categoria
	if err := s.db.First(&existingCategoria, id).Error; err != nil {
		return err
	}
	return s.db.Model(&existingCategoria).Updates(&categoria).Error
}

func (s *categoriaServiceImpl) DeleteCategoria(id int) error {
	return s.db.Delete(&models.Categoria{}, id).Error
}

// Handler functions

func GetCategorias(c *gin.Context, service CategoriaService) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	categorias, err := service.GetCategorias(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categorias)
}

func GetCategoriaByID(c *gin.Context, service CategoriaService) {
	id := c.Param("id")
	categoria, err := service.GetCategoriaByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func PostCategoria(c *gin.Context, service CategoriaService) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateCategoria(categoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, categoria)
}

func UpdateCategoria(c *gin.Context, service CategoriaService) {
	id := c.Param("id")
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.UpdateCategoria(id, categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func DeleteCategoria(c *gin.Context, service CategoriaService) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.DeleteCategoria(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoria deleted"})
}
