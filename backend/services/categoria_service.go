package services

import (
	"net/http"
	"robotica_concursos/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategorias(c *gin.Context, page int, pageSize int) {
	var categorias []models.Categoria
	// Logic to get all categorias
	var db = GetDatabase()
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&categorias)
	c.JSON(http.StatusOK, categorias)
}

func GetCategoriaByID(c *gin.Context) {
	id := c.Param("id")
	var categoria models.Categoria
	// Logic to get categoria by ID
	var db = GetDatabase()
	result := db.First(&categoria, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func PostCategoria(c *gin.Context) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to create a new categoria
	var db = GetDatabase()
	if err := db.Create(&categoria).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, categoria)
}

func UpdateCategoria(c *gin.Context) {
	id := c.Param("id")
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to update categoria by ID
	var db = GetDatabase()
	result := db.First(&models.Categoria{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id no encontrado"})
		return
	}
	if err := db.Model(&categoria).Updates(&categoria).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func DeleteCategoria(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to delete categoria by ID
	var db = GetDatabase()
	var categoria models.Categoria
	categoria.ID = uint(id)
	db.Delete(&categoria)
	c.JSON(http.StatusOK, gin.H{"message": "Categoria deleted"})
}
