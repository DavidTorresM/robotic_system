package controllers

import (
	"net/http"

	"robotica_concursos/models"
	"robotica_concursos/services"

	"github.com/gin-gonic/gin"
)

type CategoriaController struct{}

func (cc CategoriaController) GetCategorias(c *gin.Context) {
	var categorias []models.Categoria
	if err := services.GetDatabase().Find(&categorias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categorias)
}

func (cc CategoriaController) GetCategoria(c *gin.Context) {
	id := c.Param("id")
	var categoria models.Categoria
	if err := services.GetDatabase().First(&categoria, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria not found"})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func (cc CategoriaController) CreateCategoria(c *gin.Context) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.GetDatabase().Create(&categoria).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, categoria)
}

func (cc CategoriaController) UpdateCategoria(c *gin.Context) {
	id := c.Param("id")
	var categoria models.Categoria
	if err := services.GetDatabase().First(&categoria, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria not found"})
		return
	}
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.GetDatabase().Save(&categoria).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func (cc CategoriaController) DeleteCategoria(c *gin.Context) {
	id := c.Param("id")
	if err := services.GetDatabase().Delete(&models.Categoria{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func RegisterCategoriaRoutes(router *gin.Engine) {
	cc := CategoriaController{}
	router.GET("/categorias", cc.GetCategorias)
	router.GET("/categorias/:id", cc.GetCategoria)
	router.POST("/categorias", cc.CreateCategoria)
	router.PUT("/categorias/:id", cc.UpdateCategoria)
	router.DELETE("/categorias/:id", cc.DeleteCategoria)
}
